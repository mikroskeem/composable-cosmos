# syntax=docker/dockerfile:1

ARG GO_VERSION="1.19"
ARG RUNNER_IMAGE="gcr.io/distroless/static-debian11"

# --------------------------------------------------------
# Builder
# --------------------------------------------------------

FROM golang:${GO_VERSION}-alpine as builder

ARG GIT_VERSION
ARG GIT_COMMIT

RUN apk add --no-cache \
    ca-certificates \
    build-base \
    linux-headers

# Download go dependencies
WORKDIR /centauri
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/root/go/pkg/mod \
    go mod download

# Cosmwasm - Download correct libwasmvm version
RUN set -eux; \    
    export ARCH=$(uname -m); \
    WASM_VERSION=$(go list -m all | grep github.com/CosmWasm/wasmvm | awk '{print $2}'); \
    if [ ! -z "${WASM_VERSION}" ]; then \
      wget -O /lib/libwasmvm_muslc.a https://github.com/CosmWasm/wasmvm/releases/download/${WASM_VERSION}/libwasmvm_muslc.${ARCH}.a; \      
    fi; \
    go mod download;
    
# Copy the remaining files
COPY . .

# Build centaurid binary
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/root/go/pkg/mod \
    GOWORK=off go build \
        -mod=readonly \
        -tags "netgo,ledger,muslc" \
        -ldflags \
            "-X github.com/cosmos/cosmos-sdk/version.Name="centauri" \
            -X github.com/cosmos/cosmos-sdk/version.AppName="centaurid" \
            -X github.com/cosmos/cosmos-sdk/version.Version=${GIT_VERSION} \
            -X github.com/cosmos/cosmos-sdk/version.Commit=${GIT_COMMIT} \
            -X github.com/cosmos/cosmos-sdk/version.BuildTags=netgo,ledger,muslc \
            -w -s -linkmode=external -extldflags '-Wl,-z,muldefs -static'" \
        -trimpath \
        -o /centauri/build/centaurid \
        /centauri/cmd/centaurid


# --------------------------------------------------------
# toolkit
# --------------------------------------------------------

FROM busybox:1.35.0-uclibc as busybox
RUN addgroup --gid 1025 -S composable && adduser --uid 1025 -S composable -G composable


# --------------------------------------------------------
# Runner
# --------------------------------------------------------
FROM ${RUNNER_IMAGE}

COPY --from=busybox:1.35.0-uclibc /bin/sh /bin/sh

COPY --from=builder /centauri/build/centaurid /bin/centaurid

# Install composable user
COPY --from=busybox /etc/passwd /etc/passwd
COPY --from=busybox --chown=1025:1025 /home/composable /home/composable

WORKDIR /home/composable
USER composable

# rest server
EXPOSE 1317
# tendermint p2p
EXPOSE 26656
# tendermint rpc
EXPOSE 26657
# grpc
EXPOSE 9090

ENTRYPOINT ["centaurid"]

#
# Build from the root of the project with the command
# docker build . -f build/package/Dockerfile -t uxxu/kuboxy
#

############################
# STEP 1 build executable binary
############################
FROM golang:1.12-rc-alpine AS build-env

# Give a label, so that this temporary image will apprear nicely in the docker list
LABEL image=builder-environment

# Add
# * git so that go can download dependencies
# * build-base to have basic build tools as we build with CGO_ENABLED=1
# * musl-dev musl-utils to build against musl
# * upx for compressing the generated executable
# And then clean the cache to reduce Docker image
RUN apk add --update-cache git build-base musl musl-dev musl-utils upx \
 && rm -rf /var/cache/apk/*

# Define where the sources are being copied in the build environment
WORKDIR /go/src/github.com/twuillemin/kuboxy

# Copy from the root of the project to destination
# All needed folders should be copied, that's cleaner than copying everything
COPY go.mod go.mod
COPY cmd cmd
COPY internal internal
COPY pkg pkg
COPY docs docs

# Define the environmnet variable for building with GO
# * Dependencies managed with mod files
# * SQLite requires CGO_ENABLED=1, which in turn require gcc which by default will use glibc
# * Target is linux 64 bits
ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64

# Build the application
# * Download dependencies of the projects and clean mod file
# * Build the application in a file main.exe that is to be stored in the root
# * Strip and compress the binary
RUN go mod download \
 && go mod tidy \
 && go mod verify \
 && go build -a -tags netgo -ldflags "-w -linkmode external -extldflags '-static' " -o kuboxy.exe cmd/main.go \
 && strip --strip-unneeded kuboxy.exe \
 && upx kuboxy.exe


############################
# STEP 2 build a small image
############################
FROM scratch

# Copy our static executable.
COPY --from=build-env /go/src/github.com/twuillemin/kuboxy/kuboxy.exe /go/bin/kuboxy.exe

# Run the hello binary.
ENTRYPOINT ["/go/bin/kuboxy.exe"]
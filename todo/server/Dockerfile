# This is supposed to be build with the BuildKit engine

FROM --platform=$BUILDPLATFORM alpine as protoc
ARG BUILDPLATFORM TARGETOS TARGETARCH

# downlaod the protoc binary from github
# e.g., https://github.com/protocolbuffers/protobuf/releases/download/v25.4/protoc-25.4-linux-x86_64.zip
RUN export PROTOC_VERSION=25.4 \
    && export PROTOC_ARCH=$(uname -m) \
    && export PROTOC_OS=$(echo $TARGETOS) \
    && export PROTOC_ZIP=protoc-$PROTOC_VERSION-$PROTOC_OS-$PROTOC_ARCH.zip \
    && echo "downloading: " https://github.com/protocolbuffers/protobuf/releases/download/v$PROTOC_VERSION/$PROTOC_ZIP \
    && wget https://github.com/protocolbuffers/protobuf/releases/download/v$PROTOC_VERSION/$PROTOC_ZIP \
    && unzip -o $PROTOC_ZIP -d /usr/local bin/protoc 'include/*' \
    && rm -f $PROTOC_ZIP

FROM --platform=$BUILDPLATFORM golang:1.22-alpine as build
ARG BUILDPLATFORM TARGETOS TARGETARCH

# copy the protoc binary and the protobuf includes
COPY --from=protoc /usr/local/bin/protoc /usr/local/bin/protoc
COPY --from=protoc /usr/local/include/google /usr/local/include/google

# download protoc plugins
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# copy proto files into go/src/proto
WORKDIR /go/src/proto
COPY ./proto .

# generate code out of proto files
WORKDIR /go/src
ENV MODULE=github.com/Nu1LSpaxe/grpc-go-for-professionals
RUN protoc --go_out=./proto \
    --go_opt=module=$MODULE \
    --go-grpc_out=./proto \
    --go-grpc_opt=module=$MODULE \
    ./proto/v2/*.proto

# copy code into go/src/app
WORKDIR /go/src/app
COPY ./server .

# download dependencies and build
RUN go mod download
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -ldflags="-s -w" -o /go/bin/app


FROM alpine:latest

# copy the previously built binary into smaller image
COPY --from=build /go/bin/app /app
EXPOSE 50051
CMD ["/app", "0.0.0.0:50051"]
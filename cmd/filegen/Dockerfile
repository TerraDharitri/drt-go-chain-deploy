
FROM golang:1.13.6 as builder
MAINTAINER Dharitri

WORKDIR /go/drt-chain-deploy-go
COPY . .
WORKDIR /go/drt-chain-deploy-go/cmd/filegen
RUN GO111MODULE=on go mod vendor
RUN go build

# Second stage
FROM ubuntu:18.04
COPY --from=builder "/go/drt-chain-deploy-go/cmd/filegen" "/go/drt-chain-deploy-go/cmd/filegen"
WORKDIR /data
ENTRYPOINT ["/go/drt-chain-deploy-go/cmd/filegen/filegen"]

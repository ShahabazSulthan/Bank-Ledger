FROM golang:1.22.0-alpine AS build

WORKDIR /project/bankingledger/

COPY go.mod go.sum ./

COPY . .

RUN go build -o ./cmd/Api_GatewayExec ./cmd/main.go

FROM alpine:latest

WORKDIR /project/bankingledger/

COPY --from=build /project/bankingledger/cmd/Api_GatewayExec ./cmd/

COPY --from=build /project/bankingledger/dev.env ./dev.env

RUN apk update && apk add --no-cache ca-certificates tzdata

EXPOSE 8000

ENTRYPOINT ["/project/bankingledger/cmd/Api_GatewayExec"]
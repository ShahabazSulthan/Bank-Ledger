FROM golang:1.23.0-alpine AS build
WORKDIR /project/bankingledger/

COPY go.mod go.sum ./
COPY . .

RUN go build -o ./cmd/Bank_ServiceExec ./cmd/main.go

FROM alpine:latest
WORKDIR /project/friendzy/

COPY --from=build /project/bankingledger/cmd/Bank_ServiceExec ./cmd/Bank_ServiceExec
COPY --from=build /project/bankingledger/dev.env ./dev.env

RUN apk update && apk add --no-cache ca-certificates tzdata

EXPOSE 50001
ENTRYPOINT ["/project/friendzy/cmd/Bank_ServiceExec"]

FROM golang:1.23.0-alpine AS build
WORKDIR /project/bankingledger/

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o ./cmd/Notif_ServiceExec ./cmd/main.go

FROM alpine:latest
WORKDIR /project/friendzy/

COPY --from=build /project/bankingledger/cmd/Notif_ServiceExec ./cmd/Notif_ServiceExec
COPY --from=build /project/bankingledger/dev.env ./dev.env

RUN apk update && apk add --no-cache ca-certificates tzdata

EXPOSE 50002

ENTRYPOINT ["./cmd/Notif_ServiceExec"]

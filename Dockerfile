FROM golang:1.20-bookworm AS builder
WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download -x

COPY . ./
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o /bin/wallet .

FROM gcr.io/distroless/base-debian12
WORKDIR /app

COPY --from=builder /bin/wallet ./

ENTRYPOINT ["./wallet"]
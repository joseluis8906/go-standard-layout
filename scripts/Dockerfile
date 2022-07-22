FROM golang:1.18-alpine AS builder
WORKDIR /app
COPY . .
RUN apk --no-cache add ca-certificates build-base pkgconf
RUN GO111MODULE=on GOS=linux go build -mod=vendor -o ./bin/go-standard-layout ./cmd/http/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY ./configs ./configs
COPY --from=builder /app/bin/go-standard-layout
CMD ["./go-standard-layout"]

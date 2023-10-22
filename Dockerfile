FROM golang:latest AS builder

WORKDIR /root/
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /root/social ./cmd/social/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /root/social /root/
COPY ./configs/ /root/configs/

EXPOSE 8080
CMD ["./social"]
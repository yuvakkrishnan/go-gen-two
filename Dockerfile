FROM golang:1.21 AS builder
WORKDIR /app
COPY . .
RUN go build -o go-gen-two

FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/go-gen-two .
CMD ["./go-gen-two"]

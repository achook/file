FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o file

FROM alpine:latest AS runner
WORKDIR /app
COPY --from=builder /app/file /app/file
CMD ["./file"]
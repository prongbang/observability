FROM golang:1.20.4-alpine3.16 AS builder
WORKDIR /go/src/github.com/prongbang/main-service
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o main-service .

FROM alpine:3.17.3
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/prongbang/main-service/main-service ./
CMD ["./main-service"]
FROM golang:1.22.2-alpine as builder

WORKDIR /app

COPY . .

RUN apk add make

RUN make build

FROM alpine:latest as release

COPY --from=builder /app/out/main /app/main

CMD ["/app/main"]

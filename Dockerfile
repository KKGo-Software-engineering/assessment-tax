FROM golang:1.21.4-alpine as build-base

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go test -v

RUN go build -o ./out/assessment-tax .


FROM alpine:3.16.2
COPY --from=build-base /app/out/assessment-tax /app/assessment-tax

ENV PORT=8080
ENV DATABASE_URL="host=host.docker.internal port=5432 user=postgres password=postgres dbname=ktaxes sslmode=disable"

EXPOSE 8080

CMD ["/app/assessment-tax"]
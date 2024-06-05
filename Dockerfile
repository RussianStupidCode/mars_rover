FROM golang:1.22 as build

WORKDIR /app

COPY . ./

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build ./cmd/main.go

# TODO можно добавить запуск тестов сюда

FROM alpine:latest as production
COPY --from=build /app/main .
CMD ["./main"]
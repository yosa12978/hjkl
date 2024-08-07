FROM golang:1.22-alpine3.20 AS builder

WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o bin/hjkl ./main.go

FROM alpine:3.20

WORKDIR /app
COPY --from=builder /app/bin .
COPY --from=builder /app/config.yaml .

RUN apk --update --no-cache add curl

EXPOSE 80
ENTRYPOINT ["./hjkl"]
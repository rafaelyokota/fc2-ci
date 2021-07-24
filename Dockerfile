FROM golang:1.13.0-alpine AS build

WORKDIR /app

COPY . .

RUN go build -o math
CMD ["./math"]

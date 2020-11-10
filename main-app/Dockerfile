FROM golang:latest as builder

ENV G0111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

FROM ubuntu:latest
COPY --from=builder /app/app /app/
EXPOSE 8000
ENTRYPOINT [ "/app/app" ]
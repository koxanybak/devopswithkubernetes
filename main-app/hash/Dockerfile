FROM golang:latest as builder

ENV G0111MODULE=on

WORKDIR /hash

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o hash

FROM ubuntu:latest
COPY --from=builder /hash/ /app/hash/
WORKDIR /app/hash
RUN mkdir ../files && touch ../files/stamp.txt 
EXPOSE 8000
ENTRYPOINT [ "./hash" ]
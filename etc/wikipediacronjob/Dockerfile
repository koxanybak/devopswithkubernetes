FROM ubuntu:latest

RUN apt-get update && apt-get install -y curl ca-certificates

COPY job.sh job.sh

ENTRYPOINT [ "bash", "job.sh" ]
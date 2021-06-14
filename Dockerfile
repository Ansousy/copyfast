FROM golang:1.16.5

COPY / /tmp/

WORKDIR /tmp

ENTRYPOINT [ "/bin/sh" ]
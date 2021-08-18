FROM alpine:latest
ENTRYPOINT ["/usr/bin/ddnss"]
COPY ddnss /usr/bin/ddnss
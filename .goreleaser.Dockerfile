FROM alpine:latest
COPY ddnss /usr/bin/ddnss
ENTRYPOINT ["/usr/bin/ddnss"]
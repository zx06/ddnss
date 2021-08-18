FROM alpine:latest
ENTRYPOINT ["/usr/bin/app"]
COPY app /usr/bin/app
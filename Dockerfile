#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/ddnss
COPY . .
RUN go mod tidy
RUN go build -o /go/bin/ddnss -v .
RUN ls -lah /go/bin/ddnss

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/ddnss /ddnss
ENTRYPOINT /ddnss
LABEL Name=ddnss Version=0.0.1
EXPOSE 5353

FROM alpine:3.21

COPY am /

ENTRYPOINT ["/am"]

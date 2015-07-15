FROM alpine

ADD cbdproxy_linux /bin/cbdproxy

ENV PORT 80
ENTRYPOINT ["/bin/cbdproxy"]

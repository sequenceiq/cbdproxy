FROM alpine

ADD https://github.com/sequenceiq/cbdproxy/releases/download/v0.0.6/cbdproxy_linux /bin/cbdproxy

ENV PORT 80
ENTRYPOINT ["/bin/cbdproxy"]

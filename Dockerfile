FROM alpine

ADD https://github.com/sequenceiq/cbdproxy/releases/download/v0.0.8/cbdproxy_linux /bin/cbdproxy

ENV PORT 80
ENTRYPOINT ["/bin/cbdproxy"]

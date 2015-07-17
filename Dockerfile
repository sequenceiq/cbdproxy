FROM alpine

ADD https://github.com/sequenceiq/cbdproxy/releases/download/vX.X.X/cbdproxy_linux /bin/cbdproxy

ENV PORT 80
ENTRYPOINT ["/bin/cbdproxy"]

DOCKER_OPTS="--tlsverify=false"
VERSION=0.0.3

build:
	GOOS=linux go build -o cbdproxy_linux

docker: build
	docker $(DOCKER_OPTS) build -t sequenceiq/cbdproxy .

test:
	docker rm -f cbdproxy 2> /dev/null || true
	docker $(DOCKER_OPTS) run -it \
		--name=cbdproxy \
		-e DEBUG=1 \
		-p 80:80 \
		--dns 192.168.59.103 \
		sequenceiq/cbdproxy


release: build
	gh-release create sequenceiq/cbdproxy $(VERSION)
	dockerhub-tag create  sequenceiq/cbdproxy $(VERSION) $(VERSION) .

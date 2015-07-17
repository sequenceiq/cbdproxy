DOCKER_OPTS="--tlsverify=false"

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
		sequenceiq/cbdproxy_linux

prepare_release:
	./generate_new_version.sh

release: build
	VERSION=$(shell cat VERSION)

	echo mingya megy ki: $(NEW_VERSION)
	#gh-release create sequenceiq/cbdproxy $(VERSION)
	#dockerhub-tag create  sequenceiq/cbdproxy $(VERSION) v$(VERSION) .

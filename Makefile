DOCKER_OPTS="--tlsverify=false"

build:
	GOOS=linux go build -o cbdproxy_linux

docker: build
	docker $(DOCKER_OPTS) build -t sequenceiq/cbdproxy .

test:
	docker $(DOCKER_OPTS) run -it \
		-p 80:80 \
		--dns 192.168.59.103 \
		sequenceiq/cbdproxy




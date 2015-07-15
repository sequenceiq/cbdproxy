DOCKER_OPTS="--tlsverify=false"

build:
	GOOS=linux go build -o cbdproxy_linux

docker: build
	docker $(DOCKER_OPTS) build -t sequenceiq/cbdproxy .

test:
	docker rm -f cbdproxy
	docker $(DOCKER_OPTS) run -it \
		--name=cbdproxy \
		-e DEBUG=1 \
		-p 80:80 \
		--dns 192.168.59.103 \
		sequenceiq/cbdproxy




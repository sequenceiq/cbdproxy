# Cloudbreak Proxy

Cloudbreak Proxy is a golang ReverseProxy in front of Cloudbreak components:
 - Cloudbreak API
 - Uluwatu (Cloudbreak UI)
 - Sultans (Cloudbreak User management)
 - UAA (OAuth2 server implementation)

Cloudbreak Proxy uses Consul provided dns SRV entries to discover service urls.

## Usage

Cloudbreak Proxy is meant to be used as a Docker container. When you use 
Cloudbreak Deployer (cbd), than Cloudbreak Proxy is managed for you.

## Manual start

For troubleshooting purpose, you might want to start it manually,
here is how to start it manually:

```
    docker run -d \
    --name cbdproxy \
    -p 80:80 \
    --dns 192.168.59.103 \
    sequenceiq/cbdproxy
```

- `-p 80:80` defines the port where the Cloudbreak Proxy listens
- `--dns 192.168.59.103` point to Consul’s bind ip

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

## Configuration

Service host:ports are discovered by asking SRV records from dns.
When used as a docker container, you set the DNS server as `--dns`

If you need an alternative way you can use the `DNS_HOST` and `DNS_PORT`
environment variables.

By default proxy listens on port 80. You can change it by the `PORT`
environment variable.

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

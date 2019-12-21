# Api Service

This is the Api service

Generated with

```
micro new api-gateway --namespace=com.lcb123 --type=web
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: com.lcb123.web.api
- Type: web
- Alias: api

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./api-web
```

Build a docker image
```
make docker
```
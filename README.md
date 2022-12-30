# memcache

Fast, simple, in-memory key-value caching using HTTP/HTTPS.

Note: All values are stored as strings in utf-8 encoding.

## Installation

### Download binary from GitHub

Go to [https://github.com/eminmuhammadi/memcache/releases/latest](https://github.com/eminmuhammadi/memcache/releases/latest) page and download the binary for your operating system.

### Install from source

```bash
git clone https://github.com/eminmuhammadi/memcache && cd memcache && chmod +x ./install.sh && ./install.sh
```

Note: GCC is required to build from source.

## Simple Usage

```bash
memcache start --hostname 0.0.0.0 --port 8080
```

## Options

```bash
   --hostname value         network interface to listen on
   --port value             network port to listen on
   --timezone value         timezone to use for time.Time (default: "UTC")
   --logLevel value         log level to use; one of: info, warn, error, silent (default: "SILENT")
   --bodyLimit value        in bytes (1024 * 1024 = 1MB) (default: 4194304)
   --readBufferSize value   in bytes (default: 4096)
   --writeBufferSize value  in bytes (default: 4096)
   --readTimeout value      in seconds (default: 15)
   --writeTimeout value     in seconds (default: 15)
   --idleTimeout value      in seconds (default: 60)
   --prefork                enable preforking
   --concurrency value      number of concurrent connections to handle (default: 262144)
   --secure                 enable TLS
   --certFile value         path to TLS certificate file
   --keyFile value          path to TLS key file
```

To see details about the available options, run `memcache start --help`. If you want to disable banner, you need to set `MEMCACHE_DISABLE_BANNER` environment variable to any value.

Note: If you want to use prefork mode in docker container, you need to run it with `CMD ./memcache` or `CMD ["sh", "-c", "/memcache"]`.

## Usage

### Set a value

```bash
curl -L -X POST '127.0.0.1:8080/' \
-H 'Content-Type: application/json' \
--data-raw '{
    "value": "Hello World!"
}'
```

### Get a value

```bash
curl -L -X GET '127.0.0.1:8080/068ddfad-9288-4c39-90b2-98e51b9759da' \
-H 'Content-Type: application/json'
```

### Delete a value

```bash
curl -L -X DELETE '127.0.0.1:8080/068ddfad-9288-4c39-90b2-98e51b9759da' \
-H 'Content-Type: application/json'
```

### Edit a value

```bash
curl -L -X PUT '127.0.0.1:8080/45620163-14d5-49c9-9e64-76f97006efea' \
-H 'Content-Type: application/json' \
--data-raw '{
    "value": "Hello New World!"
}'
```

## Metrics

Following metrics are available:

- `<protocol>://<ip>:<port>/_/metrics` - Prometheus metrics
- `<protocol>://<ip>:<port>/_/healthcheck` - Health check
- `<protocol>://<ip>:<port>/_/monitoring` - Monitoring (OS Information, Memory, CPU, etc.)

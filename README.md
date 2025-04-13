# Public IP Address Change Notifier

A simple app that sends Discord notifications whenever the public IP address changes.

## Prerequisites
- Go 1.24 or later.
- Docker (optional).

## Installation

Before running ensure that all required envs are set.

### Environment Variables

- `DISCORD_WEBHOOK_URL`. Required. The webhook URL for sending Discord notifications.
- `SCHEDULE`. Optional. Default is `@every 1h`. A CRON string supported by [robfig/cron/v3](https://pkg.go.dev/github.com/robfig/cron/v3#hdr-Usage).
- `LOG_LEVEL`. Optional. Default is `info`. The logging level for the application (`error`, `info`, `debug`, etc.). 

### Running Locally
   ```bash
   go run .
   ```

### Building Locally
1. Build the application:
   ```bash
   go build -o ip-addr-notifier .
   ```
2. Run the application:
   ```bash
   ./ip-addr-notifier
   ```

### Using Docker
1. Build the Docker image:
   ```bash
   docker build -t ip-addr-notifier .
   ```
2. Run the Docker container:
   ```bash
   docker run --rm ip-addr-notifier
   ```

### Using Docker Compose

```yaml
services:
  ip-addr-notifier:
    image: ghcr.io/exec-astraea/ip-change-notifier:latest
    environment:
      - DISCORD_WEBHOOK_URL=${DISCORD_WEBHOOK_URL}
      - SCHEDULE=${SCHEDULE}
    volumes:
      - /srv/ip-addr-notifier/last_ip.txt:/app/last_ip.txt
```

## License

This project is licensed under the GNU General Public License v3.0. See the [LICENSE](./LICENSE) file for details.
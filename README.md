# tyk-app

A demo application showing a Node.js backend proxied through the [Tyk API Gateway](https://tyk.io), with JWT authentication and custom Go middleware plugins deployed via Jenkins CI.

## Architecture

```
Client → Tyk Gateway (/dummy-proxy/) → dummy-backend:3001
                ↕
         Tyk Dashboard (management)
                ↕
         PostgreSQL (analytics)
```

**Go plugins** run inline in the gateway:
- `add_channel` — pre-plugin: injects a `X-Channel` request header
- `transform_response` — response plugin: transforms the upstream response body

## Project Structure

```
tyk-app/
├── backend/                        # Node.js upstream service
│   ├── server.js
│   └── Dockerfile
│
├── tyk-config/                     # All Tyk configuration
│   ├── apis/
│   │   └── dummy/
│   │       └── dummy-api.json      # OAS 3.0 API definition (x-tyk-api-gateway)
│   ├── policies/
│   │   └── dummy/
│   │       └── dummy-api-policy.json
│   ├── middleware/
│   │   └── dummy/
│   │       └── golang/             # Go plugin source (compiled by CI)
│   │           ├── add_channel/
│   │           └── transform_response/
│   ├── assets/
│   │   └── templates/              # Error templates, transform templates
│   ├── environments/
│   │   └── dummy/                  # dev.env, staging.env, prod.env
│   ├── scripts/                    # Utility scripts (e.g. validate-apis.sh)
│   └── gateway/
│       ├── tyk.json                # Gateway configuration
│       └── Dockerfile
│
├── Jenkinsfile.tyk                 # Tyk deploy pipeline
├── Jenkinsfile.backend             # Backend deploy pipeline
└── docker-compose.yml              # Local development stack
```

## Prerequisites

- Docker & Docker Compose
- Jenkins with Docker access (socket mounted)
- Jenkins credential: `tyk-gw-secret` (Tyk gateway secret)

## Local Development

Start the full stack (Gateway, Dashboard, PostgreSQL, Redis, backend):

```bash
docker compose up -d
```

Get the initial Jenkins admin password:

```bash
docker exec jenkins cat /var/jenkins_home/secrets/initialAdminPassword
```

Get the Tyk Dashboard admin credentials (after first login):

```bash
docker exec tyk-postgres psql -U postgres -d tyk_analytics \
  -t -c "SELECT accesskey, orgid FROM tyk_analytics_users WHERE emailaddress='dev@tyk.io' LIMIT 1;"
```

## Jenkins Pipeline (`Jenkinsfile.tyk`)

The pipeline runs the following stages in order:

| Stage | Description |
|---|---|
| **Checkout** | Pulls latest from `main` |
| **Get Dashboard Credentials** | Fetches live API key + org ID from PostgreSQL |
| **Build and Deploy Go Plugins** | Compiles `.so` files with `tyk-plugin-compiler` and copies to `tyk_mw` volume |
| **Delete API from Dashboard** | Removes old API definition (idempotent) |
| **Delete Policy from Dashboard** | Removes old policy (idempotent) |
| **Create Policy** | Posts `dummy-api-policy.json` to Dashboard |
| **Create API in Dashboard** | Posts `dummy-api.json` to OAS endpoint |
| **Reload Gateway** | Triggers hot reload via `GET /tyk/reload/group` |

## Go Plugins

Plugins are compiled using the official Tyk plugin compiler to ensure ABI compatibility:

```bash
docker run --rm \
  -e PLUGIN_SOURCE_PATH=/path/to/plugin \
  tykio/tyk-plugin-compiler:v5.11.0 plugin_name.so "" linux arm64
```

> Compiled `.so` binaries are **not committed to git** — they are built fresh in CI on every run.

## Adding a New Service

1. Create `tyk-config/apis/<service>/` with an OAS API definition
2. Create `tyk-config/policies/<service>/` with a policy JSON
3. Add Go plugins under `tyk-config/middleware/<service>/golang/<plugin-name>/`
4. Add a corresponding stage in `Jenkinsfile.tyk`

# Goblin Wrangler
Welcome to Goblin Wrangler, a tabletop gaming service!

## Local development
Prior to running Goblin Wrangler locally you must do some setup:

1. `cp web/.env.development.example web/.env.development` and set the GOOGLE_CLIENT_ID to a valid OAuth 2 Client ID you've generated in GCP.

After that it is possible to deploy Goblin Wrangler locally using [Docker Compose](https://docs.docker.com/compose/). Our compose file will deploy the 
Goblin Wrangler services and required databases from scratch. To deploy the service, run:

```bash
docker-compose build
docker-compose up -d
```

The website will then be available via [localhost](http://localhost:3000). Both the backend and web projects reload on code changes.

## Remote deployment

Goblin Wrangler currently makes use of the following:
* DNS via CloudFlare.
* OAuth credentials from Google GCP.
* Netlify to build and publish the frontend project (`./web`), a Vite/Vue3 project.
* Fly.io to build and publish the backend project (`./backend`), a Golang project.
* Fly.io for the Postgres DB used by the backend.

Netlify stores secrets related to the frontend, and Fly.io secrets related to the backend. The secrets are not currently codified
or stored in a secrets manager of any kind, and the infrastructure and setup on Netlify and Fly are not codified beyond the relevant
deployment TOML files in each project.
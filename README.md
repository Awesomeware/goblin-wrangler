# Goblin Wrangler
Welcome to Goblin Wrangler, a tabletop gaming service!

## Local development
Prior to running Goblin Wrangler locally you must do some setup:

1. `cp web/.firebaserc.example web/.firebaserc` and change the my-gcp-project text to point to the GCP project you've attached to Firebase Hosting.
1. `cp web/.env.development.example web/.env.development` and set the GOOGLE_CLIENT_ID to a valid OAuth 2 Client ID you've generated in GCP.

After that it is possible to deploy Goblin Wrangler locally using [Docker Compose](https://docs.docker.com/compose/). Our compose file will deploy the 
Goblin Wrangler services and required databases from scratch. To deploy the service, run:

```bash
docker-compose build
docker-compose up -d
```

The website will then be available via [localhost](http://localhost:3000). Both the backend and web projects reload on code changes.

## Project components

This project consists of a Go-based backend (found under `./backend`), and a Vite/Vue3 TypeScript frontend (found under `./web`).

## Deployment

Goblin Wrangler makes use of Google services to deploy:
* Cloud Run for the backend, with a cloud run trigger to build new releases of backend/ on repository changes.
* Firebase Hosting for frontend, using GitHub CI/CD to build new releases of web/ on repository changes.

The infrastructure to support this is not currently codified, but the idea is to do so through e.g., [pulumi](https://www.pulumi.com/)
or [terraform](https://www.terraform.io/) soon, in a way that would let you deploy remotely in the same way. But you could deploy
the backend and frontend in any way you like (such as EKS, via Vercel, etc), without any real changes to the project parts.
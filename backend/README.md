# Various tidbids

Project structure based on https://github.com/golang-standards/project-layout

Designed to be transitioned to a microservice architecture over time, but I couldn't
really be bothered to deploy a ton of services and thus need to dip into k8s straight
away. So for now the whole backend is contained in 'backend' and 'web' for a backend
service and web SPA respectively.

Currently build into Fly.io. The following secrets need to exist:

* VPR_DATABASE_URL: a valid postgres:// database URL.
* VPR_FRONTEND_CORS_ORIGIN: a CORS origin to set for the web server.
* VPR_GOOGLE_CLIENT_ID: a valid Google client ID for SSO-based OAuth flow.

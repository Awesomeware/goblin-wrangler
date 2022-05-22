# Various tidbids

Project structure based on https://github.com/golang-standards/project-layout

Designed to be transitioned to a microservice architecture over time, but I couldn't
really be bothered to deploy a ton of services and thus need to dip into k8s straight
away. So for now the whole backend is contained in 'backend' and 'web' for a backend
service and web SPA respectively.
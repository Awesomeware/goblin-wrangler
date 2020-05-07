# Goblin Wrangler
Welcome to Goblin Wrangler, a tabletop gaming service!

## Local Deployment
It is possible to deploy Goblin Wrangler locally using [Docker Compose](https://docs.docker.com/compose/). Our compose file will deploy the Goblin Wrangler services and required databases from scratch. To deploy the service, run:

```bash
docker-compose build
docker-compose up -d
```

The service will then be available via [Localhost](http://localhost:4000).
# Artificer - Infrastructure for Goblin Wrangler
This folder handles all infrastructure we need to run the Goblin Wrangler. We currently operate Goblin Wrangler inside a Kubernetes cluster, deployed on [Digital Ocean](https://www.digitalocean.com/). This is provisioned with [Terraform](https://www.terraform.io/).

## Deploying Infrastructure
To deploy infrastructure from scratch you need to do the following:

```bash
terraform init
DIGITALOCEAN_TOKEN=<API key> terraform plan
DIGITALOCEAN_TOKEN=<API key> terraform apply
```

This will deploy a Kubernetes cluster in Digital Ocean, along with the following cluster-wide services:
 - Nginx Ingress Controller
 - cert-manager to provide [Lets Encrypt](https://letsencrypt.org/) certificates to TLS secure Ingresses.
 - [Jenkins](https://www.jenkins.io/) to provide CI within the cluster.
# Stuff we can change via command line
variable "cluster_name" {
    type    = string
    default = "goblin-wrangler"
}

variable "cluster_domain" {
    type    = string
    default = "goblinwrangler.com"
}

variable "cluster_region" {
    type    = string
    default = "fra1"
}

variable "cluster_version" {
    type    = string
    default = "1.16.6-do.2"
}

variable "cluster_pool_name" {
    type    = string
    default = "goblin-wrangler-pool"
}

variable "cluster_pool_spec" {
    type    = string
    default = "s-6vcpu-16gb"
}

variable "cluster_pool_count" {
    type    = number
    default = 2
}

# Providers be here
provider "digitalocean" {
    version = "~> 1.16"
}

# Resources be here
resource "digitalocean_kubernetes_cluster" "goblin-wrangler-cluster" {
    name    = var.cluster_name
    region  = var.cluster_region
    version = var.cluster_version

    node_pool {
        name       = var.cluster_pool_name
        size       = var.cluster_pool_spec
        node_count = var.cluster_pool_count
    }
}

provider "kubernetes" {
    version                = "~> 1.11"
    load_config_file       = false
    host                   = digitalocean_kubernetes_cluster.goblin-wrangler-cluster.endpoint
    token                  = digitalocean_kubernetes_cluster.goblin-wrangler-cluster.kube_config[0].token
    cluster_ca_certificate = base64decode(digitalocean_kubernetes_cluster.goblin-wrangler-cluster.kube_config[0].cluster_ca_certificate)
}

provider "kustomization" {
    kubeconfig_raw = digitalocean_kubernetes_cluster.goblin-wrangler-cluster.kube_config[0].raw_config
}

resource "kubernetes_namespace" "goblin-wrangler-ingress" {
    metadata {
        labels = {
            "app.kubernetes.io/name"    = "ingress-nginx",
            "app.kubernetes.io/part-of" = "ingress-nginx"
        }

        name = "ingress-nginx"
    }
}

resource "kubernetes_namespace" "goblin-wrangler-cert-manager" {
    metadata {
        name = "cert-manager"
    }
}

resource "kubernetes_namespace" "goblin-wrangler-ci" {
    metadata {
        name = "ci"
    }
}

data "kustomization" "goblin-wrangler-ingress" {
    path = "ingress"
}

resource "kustomization_resource" "goblin-wrangler-ingress" {
    for_each = data.kustomization.goblin-wrangler-ingress.ids

    manifest = data.kustomization.goblin-wrangler-ingress.manifests[each.value]
}

data "kustomization" "goblin-wrangler-cert-manager" {
    path = "cert-manager"
}

resource "kustomization_resource" "goblin-wrangler-cert-manager" {
    for_each = data.kustomization.goblin-wrangler-cert-manager.ids

    manifest = data.kustomization.goblin-wrangler-cert-manager.manifests[each.value]
}

data "kustomization" "goblin-wrangler-ci" {
    path = "ci"
}

resource "kustomization_resource" "goblin-wrangler-ci" {
    for_each = data.kustomization.goblin-wrangler-ci.ids

    manifest = data.kustomization.goblin-wrangler-ci.manifests[each.value]
}
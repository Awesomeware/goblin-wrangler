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
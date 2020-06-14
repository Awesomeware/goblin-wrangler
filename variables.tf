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
    default = "1.17.5-do.0"
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

variable "github_username" {
    type    = string
    default = "Stephen001"
}

variable "github_access_token" {
    type    = string
}

variable "github_oauth_id" {
    type    = string
}

variable "github_oauth_secret" {
    type    = string
}
module "artificer" {
    source = "./artificer"

    cluster_name = var.cluster_name
    cluster_domain = var.cluster_domain
    cluster_region = var.cluster_region
    cluster_version = var.cluster_version
    cluster_pool_name = var.cluster_pool_name
    cluster_pool_spec = var.cluster_pool_spec
    cluster_pool_count = var.cluster_pool_count
    github_username = var.github_username
    github_access_token = var.github_access_token
    github_oauth_id = var.github_oauth_id
    github_oauth_secret = var.github_oauth_secret
}
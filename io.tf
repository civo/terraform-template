variable "civo_token" {}

variable "region" {
  type    = string
  default = "FRA1"
  description = "The region to provision the cluster against"
}

variable "cluster_name_prefix" {
  type    = string
  default = "tf-template-"
}

variable "cluster_node_size" {
  type        = string
  default     = "g4s.kube.medium"
  description = "The size of the nodes to provision. Run `civo size list` for all options"
}

variable "cluster_node_count" {
  type    = number
  default = 3
}

# Firewall Access

variable "kubernetes_api_access" {
  type    = list(any)
  default = ["0.0.0.0/0"]
}

variable "cluster_web_access" {
  type    = list(any)
  default = ["0.0.0.0/0"]
}

variable "cluster_websecure_access" {
  type    = list(any)
  default = ["0.0.0.0/0"]
}

# Output

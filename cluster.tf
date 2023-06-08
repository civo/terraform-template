resource "civo_kubernetes_cluster" "cluster" {
  name        = "${var.cluster_name_prefix}cluster"
  firewall_id = civo_firewall.firewall.id
  pools {
    node_count = var.cluster_node_count
    size       = var.cluster_node_size
    # label = "my-pool-label" # This label will be set as an annotation on the nodes in the pool
  }
  timeouts {
    create = "5m"
  }
}

resource "local_file" "cluster-config" {
  content              = civo_kubernetes_cluster.cluster.kubeconfig
  filename             = "${path.module}/kubeconfig"
  file_permission      = "0600"
  directory_permission = "0755"
}
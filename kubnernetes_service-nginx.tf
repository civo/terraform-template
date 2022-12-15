
resource "kubernetes_service" "nginx" {
  metadata {
    name      = "nginx"
    namespace = "default"
annotations = {
    "kubernetes.civo.com/loadbalancer-enable-proxy-protocol" = "send-proxy"
}

  }

  spec {

    selector = {
      nginx = "nginx"
    }
    type = "LoadBalancer"
    port {
      protocol    = "TCP"
      port        = 80
      target_port = 8081
      name        = "web"
    }
    port {
      protocol    = "TCP"
      port        = 443
      target_port = 8444
      name        = "websecure"
    }

    external_traffic_policy = "Cluster"
  }

}

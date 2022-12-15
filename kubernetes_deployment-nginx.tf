resource "kubernetes_deployment" "nginx" {
  metadata {
    name      = "nginx"
    namespace = "default"
  }

  spec {
    replicas = 1
    selector {
      match_labels = {
        nginx = "nginx"
      }
    }
    template {
      metadata {
        labels = {
          nginx = "nginx"
        }
      }

      spec {
        container {
          image = "dmajrekar/nginx-echo:latest"
          name  = "nginx"
          port {
            protocol       = "TCP"
            container_port = "8080"
          }
          port {
            protocol       = "TCP"
            container_port = "8443"
          }
          resources {
            limits = {
              cpu    = "0.5"
              memory = "512Mi"
            }
            requests = {
              cpu    = "250m"
              memory = "50Mi"
            }
          }
        }
      }
    }
  }
}

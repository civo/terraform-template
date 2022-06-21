# Create a firewall
resource "civo_firewall" "firewall-ingress" {
  name                 = "${var.cluster_name_prefix}firewall-ingress"
  create_default_rules = false
}

# Create a firewall rule
resource "civo_firewall_rule" "web" {
  firewall_id = civo_firewall.firewall-ingress.id
  protocol    = "tcp"
  start_port  = "80"
  end_port    = "80"
  cidr        = var.cluster_web_access
  label       = "web"
  action      = "allow"
  direction   = "ingress"
}

resource "civo_firewall_rule" "websecure" {
  firewall_id = civo_firewall.firewall-ingress.id
  protocol    = "tcp"
  start_port  = "443"
  end_port    = "443"
  cidr        = var.cluster_websecure_access
  label       = "websecure"
  action      = "allow"
  direction   = "ingress"
}



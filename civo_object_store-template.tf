/*
resource "civo_object_store" "template" {
  name        = "${var.cluster_name_prefix}objectstore"
  max_size_gb = var.object_store_size
}


data "civo_object_store_credential" "backup" {
    id = civo_object_store.template.access_key_id
}
*/
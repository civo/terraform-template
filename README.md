# terraform-template

Opinionated Template Repo for managing applications on a Civo Kubernetes cluster 

## HLD

This is an overview of the infrastructure being managed by this repo.  

![HLD](./hld.excalidraw.png)

## Terraform Provider Documentation

* https://registry.terraform.io/providers/civo/civo/1.0.18
* https://registry.terraform.io/providers/hashicorp/helm/2.6.0
* https://registry.terraform.io/providers/hashicorp/kubernetes/2.11.0
* https://registry.terraform.io/providers/hashicorp/local/2.2.3
## Variables

| Name | Type | Description | Default |
|------|------|-------------|---------|
| civo_token | string | API Token for civo.com | "" | 
| kuberentes_api_access |  list | list of IP addresses / subnets to allow access to the cluster api | [ "0.0.0.0/0" ] |
| cluster_web_access | list | list of IP addresses / subnets to allow access to port 80 | [ "0.0.0.0/0" ] |
| cluster_websecure_access | list | list of IP addresses / subnets to allow access to port 443 | [ "0.0.0.0/0" ] |

## Contribution Guide

> TBC

## Acknowledgements

- https://www.hashicorp.com
- https://github.com/excalidraw/excalidraw
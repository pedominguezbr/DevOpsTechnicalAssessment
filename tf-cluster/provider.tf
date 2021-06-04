variable "ibmcloud_api_key" {}
variable "environment" {}
variable "resource_group" {}


provider "ibm" {
  generation       = "1"
  region           = "us-south"
  ibmcloud_api_key = var.ibmcloud_api_key
}

data "ibm_resource_group" "resource_group" {
  name = var.resource_group
}

resource ibm_container_cluster "tfcluster" {
    name            = "${var.environment}-tfcluster"
    datacenter      = "dal10"
    machine_type    = "free"
    hardware        = "shared"
    kube_version = "1.19"

    public_service_endpoint  = "true"
    private_service_endpoint = "true"

    resource_group_id = data.ibm_resource_group.resource_group.id
}
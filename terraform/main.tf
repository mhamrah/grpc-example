variable "region" {
	default = "us-east1"
}

variable "project" {
}

variable "credentials" {
}

provider "google" {
  region = "${var.region}"
  project = "${var.project}"
  credentials = "${file(var.credentials)}"
  // Provider settings to be provided via ENV variables
}


data "google_compute_zones" "available" {
  project = "${var.project}"
}

variable "cluster_name" {
  default = "grpc-demo-cluster"
}

variable "username" {}
variable "password" {}

resource "google_container_cluster" "primary" {
  name = "${var.cluster_name}"
  region = "${var.region}"
  initial_node_count = 1
  min_master_version = "1.10" 

  master_auth {
    username = "${var.username}"
    password = "${var.password}"
  }

  node_config {
    machine_type = "f1-micro"
    oauth_scopes = [
      "https://www.googleapis.com/auth/compute",
      "https://www.googleapis.com/auth/service.management",
      "https://www.googleapis.com/auth/devstorage.read_only",
      "https://www.googleapis.com/auth/logging.write",
      "https://www.googleapis.com/auth/monitoring"
    ]
    preemptible = "true"
	  labels {
		  project = "grpc-demo"
	  }

	  tags = [ "grpc", "demo" ]
  }

}

output "cluster_name" {
  value = "${google_container_cluster.primary.name}"
}

output "endpoint" {
  value = "${google_container_cluster.primary.endpoint}"
}

output "node_version" {
  value = "${google_container_cluster.primary.node_version}"
}

# The following outputs allow authentication and connectivity to the GKE Cluster.
output "client_certificate" {
  value = "${google_container_cluster.primary.master_auth.0.client_certificate}"
}

output "client_key" {
  value = "${google_container_cluster.primary.master_auth.0.client_key}"
}

output "cluster_ca_certificate" {
  value = "${google_container_cluster.primary.master_auth.0.cluster_ca_certificate}"
}

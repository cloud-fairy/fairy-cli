# ACME Component

variable "config" {
  type        = any
}

variable "dependency" {
  type        = any
}

locals {
  networkId = try(var.dependency.vpc.vpcId, null)
  region = try(var.dependency.cloud_provider.region, null)
}

resource "some_resource" {
  region = local.region
  networkId = local.networkId
}

output "cfout" {
  value       = {
    # Displayed to the end-user as how-to's
    _instructions = [
      "Deployment instructions description here",
      "Other tips here"
    ]
    # Enable connectivity with other modules requires your interface
    foo = "foo"
    bar = "bar"
  }
}

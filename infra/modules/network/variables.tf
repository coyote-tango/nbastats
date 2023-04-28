variable "vpc_cidr" {
  description = "The CIDR block for the VPC"
  type        = string
  default     = "172.16.0.0/16"
}

variable "subneta_cidr" {
  description = "The CIDR block for Subnet 1"
  type        = string
  default     = "172.16.1.0/24"
}

variable "subnetb_cidr" {
  description = "The CIDR block for Subnet 2"
  type        = string
  default     = "172.16.2.0/24"
}

variable "subnetc_cidr" {
  description = "The CIDR block for Subnet 3"
  type        = string
  default     = "172.16.3.0/24"
}

variable "common_tags" {
  description = "A map of common tags to apply to resources"
  type        = map(string)
  default = {module = "network"}
}

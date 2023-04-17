variable "aws_region" {
  description = "aws regions"
  default     = "us-east-1"
}

variable "aws_asn" {
  description = "AWS ASN"
  default = 64512
}

variable "vpc_1_resource" {
  type = object({
    vpc_id = string,
    vpc_cidr = string,
    public_subnet_id = string,
    private_subnet_id = string,
    private_subnet_routing_table_id = string
  })
  description = "AWS VPC 1"
}

variable "vpc_2_resource" {
  type = object({
    vpc_id = string,
    vpc_cidr = string,
    public_subnet_id = string,
    private_subnet_id = string,
    private_subnet_routing_table_id = string
  })
  description = "AWS VPC 2"
}

variable "tgw_id" {
  description = "Transit Gateway Id"
}

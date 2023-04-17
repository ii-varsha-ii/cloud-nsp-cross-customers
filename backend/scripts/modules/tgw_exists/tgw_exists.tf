provider "aws" {
  region = var.aws_region
}

data "aws_ec2_transit_gateway" "tgw" {
  filter {
    name   = "options.amazon-side-asn"
    values = [var.aws_asn]
  }
}

module "main" {
  source = "../main"

  tgw_id = data.aws_ec2_transit_gateway.tgw.id
  vpc_1_resource = var.vpc_1_resource
  vpc_2_resource = var.vpc_2_resource
}
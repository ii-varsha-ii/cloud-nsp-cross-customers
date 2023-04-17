provider "aws" {
  region = var.aws_region
}

resource "aws_ec2_transit_gateway" "tgw" {
  description = "Transit gateway"
  amazon_side_asn = var.aws_asn
  tags = {
    "Name" = "TGW"
  }
}

module "main" {
  source = "../main"

  tgw_id = aws_ec2_transit_gateway.tgw.id
  vpc_1_resource = var.vpc_1_resource
  vpc_2_resource = var.vpc_2_resource
}
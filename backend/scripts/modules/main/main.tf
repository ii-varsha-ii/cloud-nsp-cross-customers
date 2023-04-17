provider "aws" {
  region = var.aws_region
}

resource "aws_ec2_transit_gateway_vpc_attachment" "tgw_attachment_to_vpc_1" {
  subnet_ids         = [var.vpc_1_resource.private_subnet_id]
  transit_gateway_id = var.tgw_id
  vpc_id             = var.vpc_1_resource.vpc_id
  tags = {
    "Name" = "Transit gateway VPC 1 attachment"
  }
}

resource "aws_ec2_transit_gateway_vpc_attachment" "tgw_attachment_to_vpc_2" {
  subnet_ids         = [var.vpc_2_resource.private_subnet_id]
  transit_gateway_id = var.tgw_id
  vpc_id             = var.vpc_2_resource.vpc_id
  tags = {
    "Name" = "Transit gateway VPC 2 attachment"
  }
}

resource "aws_route" "vpc_1_to_tgw_route" {
  route_table_id = var.vpc_1_resource.private_subnet_routing_table_id
  transit_gateway_id = var.tgw_id
  destination_cidr_block = var.vpc_2_resource.vpc_cidr
}

resource "aws_route" "vpc_2_to_tgw_route" {
  route_table_id = var.vpc_2_resource.private_subnet_routing_table_id
  transit_gateway_id = var.tgw_id
  destination_cidr_block = var.vpc_1_resource.vpc_cidr
}



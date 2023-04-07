resource "aws_vpc" "vpc" {
  cidr_block = var.vpc_cidr
  tags = merge(var.common_tags,{
    Resource = "VPC"
  
})
  
}


resource "aws_subnet" "subnet_a" {
  vpc_id            = aws_vpc.vpc.id
  cidr_block        =  var.subneta_cidr
  availability_zone = "us-west-2a"
}


resource "aws_subnet" "subnet_b" {
  vpc_id            = aws_vpc.vpc.id
  cidr_block        =  var.subnetb_cidr
  availability_zone = "us-west-2b"
}


resource "aws_subnet" "subnet_c" {
  vpc_id            = aws_vpc.vpc.id
  cidr_block        = var.subnetc_cidr
  availability_zone = "us-west-2c"
}

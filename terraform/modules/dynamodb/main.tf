terraform {
  required_version = ">= 1.7"

  required_providers {
    aws = {
        source  = "hashicorp/aws"
        version = "~> 5.0"
    }
  }
}

provider "aws" {
  region = "ap-northeast-1"
}

variable "aws_region" {
  type        = string
}

variable "env" {
  type        = string
}

variable "project_name" {
  type = string
}

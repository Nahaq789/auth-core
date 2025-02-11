resource "aws_dynamodb_table" "dynamodb" {
  name           = "${var.env}-${var.project_name}-user"
  hash_key       = "user_id"
  write_capacity = 1
  read_capacity  = 1

  attribute {
    name = "user_id"
    type = "S"
  }
  attribute {
    name = "email"
    type = "S"
  }

  global_secondary_index {
    name               = "email-index"
    hash_key           = "email"
    projection_type    = "ALL"
    write_capacity = 1
    read_capacity  = 1
  }

  tags = {
    Name = "${var.env}-${var.project_name}"
  }
}
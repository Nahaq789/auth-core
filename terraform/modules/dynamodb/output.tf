output "table_arns" {
  description = "ARNs of the DynamoDB tables"
  value = {
    users = aws_dynamodb_table.dynamodb.arn
  }
}

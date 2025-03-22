variable "dynamodb_table_arns" {
  description = "ARNs of DynamoDB tables"
  type        = map(string)
}

variable "user_table" {
  description = "dynamodb user table"
  type        = string
}

variable "client_id" {
  description = "Client ID for cognito"
  type        = string
}

variable "client_secret" {
  description = "Client secret for cognito"
  type        = string
  sensitive   = true 
}

resource "aws_iam_role" "lambda" {
  name = "${var.env}-${var.project_name}-lambda-role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "lambda.amazonaws.com"
        }
      }
    ]
  })
}

resource "aws_iam_policy" "lambda_dynamodb" {
  name        = "${var.env}-${var.project_name}-lambda-dynamodb-policy"
  description = "Allow Lambda to access DynamoDB tables"

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
          "dynamodb:PutItem",
          "dynamodb:GetItem",
          "dynamodb:UpdateItem",
          "dynamodb:DeleteItem",
          "dynamodb:Query",
          "dynamodb:Scan"
        ]
        Resource = [
          values(var.dynamodb_table_arns)[0],
          "${values(var.dynamodb_table_arns)[0]}/index/email-index"
        ] 
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "lambda_basic" {
  role       = aws_iam_role.lambda.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

resource "aws_iam_role_policy_attachment" "lambda_dynamodb" {
  role       = aws_iam_role.lambda.name
  policy_arn = aws_iam_policy.lambda_dynamodb.arn
}

resource "aws_lambda_function" "app" {
  function_name = "${var.env}-${var.project_name}"
  role          = aws_iam_role.lambda.arn
  package_type  = "Image"
  image_uri     = "${aws_ecr_repository.ecr.repository_url}:latest"

  memory_size = 128
  timeout     = 30

  environment {
    variables = {
      ENVIRONMENT = var.env
      PORT = "8080"
      USER_TABLE = var.user_table
      CLIENT_ID = var.client_id
      CLIENT_SECRET = var.client_secret
    }
  }
}

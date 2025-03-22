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

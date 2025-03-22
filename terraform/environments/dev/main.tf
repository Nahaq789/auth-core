module "lambda" {
        source = "../../modules"

        aws_region   = "ap-northeast-1"
        env = "dev"
        project_name = "auth-core"
        user_table = var.user_table
        client_id = var.client_id
        client_secret = var.client_secret
        dynamodb_table_arns = module.dynamodb.table_arns
}

module "dynamodb" {
        source = "../../modules/dynamodb"

        aws_region   = "ap-northeast-1"
        env = "dev"
        project_name = "auth-core"
}

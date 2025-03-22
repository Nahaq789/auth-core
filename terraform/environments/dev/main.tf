module "lambda" {
        source = "../../modules"

        aws_region   = "ap-northeast-1"
        env = "dev"
        project_name = "auth-core"
        user_table = var.user_table
        client_id = var.client_id
        client_secret = var.client_secret
}

module "dynamodb" {
        source = "../../modules/dynamodb"

        aws_region   = "ap-northeast-1"
        env = "dev"
        project_name = "auth-core"
}

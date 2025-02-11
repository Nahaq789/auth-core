module "lambda" {
        source = "../../modules"

        aws_region   = "ap-northeast-1"
        env = "dev"
        project_name = "auth-core"
}

module "dynamodb" {
        source = "../../modules/dynamodb"

        aws_region   = "ap-northeast-1"
        env = "dev"
        project_name = "auth-core"
}
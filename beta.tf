# beta.tf

terraform {
  required_providers {
    aws = ">= 5.36.0, <= 5.36.0"
  }
}

provider "aws" {
  region = var.AWS_REGION
}

# Lambda function
resource "aws_lambda_function" "silence_of_the_lambdas" {
  function_name = var.LAMBDA_FUNCTION_NAME
  handler      = data.aws_lambda_function.existing_lambda.handler
  runtime      = data.aws_lambda_function.existing_lambda.runtime
  role         = data.aws_lambda_function.existing_lambda.role
  filename     = var.ZIP_FILE_PATH

  source_code_hash = filebase64(var.ZIP_FILE_PATH)
}

# Fetch existing Lambda function details
data "aws_lambda_function" "existing_lambda" {
  function_name = var.LAMBDA_FUNCTION_NAME
}

# Fetch existing API Gateway
data "aws_api_gateway_rest_api" "existing_api" {
  name = "Silence of the Lambdas"
}

# Gateway Beta Deployment
resource "aws_api_gateway_deployment" "beta_deployment" {
  rest_api_id = data.aws_api_gateway_rest_api.existing_api.id
  stage_name  = "beta"
}

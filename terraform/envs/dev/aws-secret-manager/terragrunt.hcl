terraform {
  source = "../../../modules/aws-secret-manager"
}

inputs = {
  region      = "us-east-1"
  environment = "dev"
  secrets     = jsondecode(file("secrets.json"))
}

variable "region" {
  description = "AWS region"
  type        = string
  default     = "us-east-1"
}

variable "environment" {
  description = "Environment"
  type        = string
  default     = "testing"
}

variable "secrets" {
  description = "Map of secrets"
  type = map(map(string))
}

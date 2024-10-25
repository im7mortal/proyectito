output "secret_arns" {
  description = "ARNs of the deployed secrets"
  value       = {for k, v in aws_secretsmanager_secret.secret : k => v.arn}
}

data "passwordstore_secret" "secret" {
  name = "<secret-name>"
}

output "secret_contents" {
  value = data.passwordstore_secret.secret.contents
}

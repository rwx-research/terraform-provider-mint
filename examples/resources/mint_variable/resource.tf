resource "mint_variable" "example" {
  vault        = "default"
  name         = "my-variable"
  secret_value = "foobar"
}

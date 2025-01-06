terraform {
  required_providers {
    mint = {
      source = "rwx-research/mint"
    }
  }
}

provider "mint" {}

resource "mint_secret" "test" {
  vault        = "default"
  name         = "foobar"
  secret_value = "test"
}

resource "mint_variable" "test" {
  vault = "default"
  name  = "foo"
  value = "bar"
}

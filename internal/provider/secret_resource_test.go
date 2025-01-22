package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestSecretResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: providerConfig + `
resource "mint_secret" "test" {
  vault        = "terraform_provider_testing"
  name         = "test-secret"
  secret_value = "foo"
  description  = "a description"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("mint_secret.test", "vault", "terraform_provider_testing"),
					resource.TestCheckResourceAttr("mint_secret.test", "name", "test-secret"),
					resource.TestCheckResourceAttr("mint_secret.test", "secret_value", "foo"),
					resource.TestCheckResourceAttr("mint_secret.test", "description", "a description"),
				),
			},
			// Update and Read testing
			{
				Config: providerConfig + `
resource "mint_secret" "test" {
  vault        = "terraform_provider_testing"
  name         = "test-secret"
  secret_value = "bar"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("mint_secret.test", "vault", "terraform_provider_testing"),
					resource.TestCheckResourceAttr("mint_secret.test", "name", "test-secret"),
					resource.TestCheckResourceAttr("mint_secret.test", "secret_value", "bar"),
					resource.TestCheckNoResourceAttr("mint_secret.test", "description"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

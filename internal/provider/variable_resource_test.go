package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestVariableResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: providerConfig + `
resource "mint_variable" "test" {
  vault = "default"
  name  = "test-var"
  value = "foo"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("mint_variable.test", "vault", "default"),
					resource.TestCheckResourceAttr("mint_variable.test", "name", "test-var"),
					resource.TestCheckResourceAttr("mint_variable.test", "value", "foo"),
				),
			},
			// ImportState testing
			{
				ResourceName:                         "mint_variable.test",
				ImportState:                          true,
				ImportStateId:                        "default/test-var",
				ImportStateVerify:                    true,
				ImportStateVerifyIdentifierAttribute: "value",
			},
			// Update and Read testing
			{
				Config: providerConfig + `
resource "mint_variable" "test" {
  vault = "default"
  name  = "test-var"
  value = "bar"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("mint_variable.test", "vault", "default"),
					resource.TestCheckResourceAttr("mint_variable.test", "name", "test-var"),
					resource.TestCheckResourceAttr("mint_variable.test", "value", "bar"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

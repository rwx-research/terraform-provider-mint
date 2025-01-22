package main

import (
	"context"
	"log"

	"github.com/rwx-research/terraform-provider-mint/internal/provider"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

// Goreleaser will overwrite this with the correct versio number upon release
var version string = "dev"

func main() {
	opts := providerserver.ServeOpts{
		Address: "registry.terraform.io/rwx-research/mint",
	}

	if err := providerserver.Serve(context.Background(), provider.New(version), opts); err != nil {
		log.Fatal(err.Error())
	}
}

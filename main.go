package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/pritunl/terraform-provider-pritunl/pritunl"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: pritunl.Provider,
	})
}

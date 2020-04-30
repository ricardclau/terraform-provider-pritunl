package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
	"github.com/kihahu/terraform-provider-pritunl/utils"
	"github.com/pritunl/terraform-provider-pritunl/provider"
)

func main() {
	utils.OutputOpen()
	defer utils.OutputClose()

	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return provider.Provider()
		},
	})
}

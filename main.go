package main

import (
	"github.com/campisiluca/terraform-provider-rdsdataservice/rdsdataservice"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: rdsdataservice.Provider})
}

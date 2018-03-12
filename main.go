package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/frankfarrell/terrassert/terrassert"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: terrassert.Provider})
}
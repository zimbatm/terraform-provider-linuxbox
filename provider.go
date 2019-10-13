package main

import (
	"github.com/draganm/terraform-provider-linuxbox/resource/swap"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{

		ResourcesMap: map[string]*schema.Resource{
			"linuxbox_swap": swap.Resource(),
		},
	}
}
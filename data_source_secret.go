package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"os/exec"
)

func dataSourceSecretRead(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	cmd := exec.Command("pass", name)
	output, err := cmd.Output()
	if err != nil {
		return err
	}
	d.Set("data", string(output))
	return nil
}

func dataSourceSecret() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSecretRead,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

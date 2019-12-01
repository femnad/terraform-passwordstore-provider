package main

import (
	"fmt"
	"os/exec"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const (
	contentsKey = "contents"
	nameKey     = "name"
)

func dataSourceSecretRead(d *schema.ResourceData, m interface{}) error {
	name := d.Get(nameKey).(string)
	cmd := exec.Command("pass", name)
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("error reading content of secret `%s`: %w", name, err)
	}
	err = d.Set(contentsKey, string(output))
	if err != nil {
		return fmt.Errorf("cannot set key `%s` of the data source: %w", contentsKey, err)
	}
	d.SetId(name)
	return nil
}

func dataSourceSecret() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSecretRead,

		Schema: map[string]*schema.Schema{
			nameKey: &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			contentsKey: &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

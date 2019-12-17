package main

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const (
	contentsKey  = "contents"
	firstLineKey = "firstline"
	nameKey      = "name"
)

func dataSourceSecretRead(d *schema.ResourceData, m interface{}) error {
	name := d.Get(nameKey).(string)
	cmd := exec.Command("pass", name)
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("error reading content of secret `%s`: %w", name, err)
	}
	secretContent := string(output)
	err = d.Set(contentsKey, secretContent)
	if err != nil {
		return fmt.Errorf("cannot set key `%s` of the data source: %w", contentsKey, err)
	}

	lines := strings.Split(secretContent, "\n")
	if len(lines) == 0 {
		return fmt.Errorf("Splitting the secret content by newlines didn't produce any lines")
	}
	firstLine := lines[0]
	err = d.Set(firstLineKey, firstLine)
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
			firstLineKey: &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

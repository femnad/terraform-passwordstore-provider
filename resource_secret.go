package main

import (
        "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSecretCreate(d *schema.ResourceData, m interface{}) error {
        return resourceSecretRead(d, m)
}

func resourceSecretRead(d *schema.ResourceData, m interface{}) error {
        return nil
}

func resourceSecretUpdate(d *schema.ResourceData, m interface{}) error {
        return resourceSecretRead(d, m)
}

func resourceSecretDelete(d *schema.ResourceData, m interface{}) error {
        return nil
}

func resourceSecret() *schema.Resource {
        return &schema.Resource{
                Create: resourceSecretCreate,
                Read:   resourceSecretRead,
                Update: resourceSecretUpdate,
                Delete: resourceSecretDelete,

                Schema: map[string]*schema.Schema{
                        "secret": &schema.Schema{
                                Type:     schema.TypeString,
                                Required: true,
                        },
                },
        }
}


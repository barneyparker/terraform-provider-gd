package main

import (
  "github.com/hashicorp/terraform/helper/schema"
)

func Provider() *schema.Provider {
  return &schema.Provider{
    Schema: map[string]*schema.Schema{
      "key": {
        Type:        schema.TypeString,
        Required:    true,
        DefaultFunc: schema.EnvDefaultFunc("GODADDY_API_KEY", nil),
        Description: "GoDaddy API Key.",
      },

      "secret": {
        Type:        schema.TypeString,
        Required:    true,
        DefaultFunc: schema.EnvDefaultFunc("GODADDY_API_SECRET", nil),
        Description: "GoDaddy API Secret.",
      },
    },

    ResourcesMap: map[string]*schema.Resource{
      "gd_record": resourceRecord(),
    },

    ConfigureFunc: providerConfigure,
  }
}

// provider config
type GoDaddyConfig struct {
  key     string
  secret  string
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
  return &GoDaddyConfig{
    key:     d.Get("key").(string),
    secret:  d.Get("secret").(string),
  }, nil
}

package main

import (
  "log"
  "github.com/hashicorp/terraform/helper/schema"
)

func resourceRecord() *schema.Resource {
  return &schema.Resource{
    Create: resourceRecordCreate,
    Read:   resourceRecordRead,
    Update: resourceRecordCreate,
    Delete: resourceRecordDelete,

    Schema: map[string]*schema.Schema{
      "domain": {
        Type: schema.TypeString,
        Required: true,
      },
      "type": {
        Type: schema.TypeString,
        Required: true,
      },
      "name": {
        Type: schema.TypeString,
        Required: true,
      },
      "data": {
        Type: schema.TypeString,
        Required: true,
      },
      "ttl": {
        Type: schema.TypeInt,
        Optional: true,
      },
      "priority": {
        Type: schema.TypeInt,
        Optional: true,
      },
    },
  }
}

func resourceRecordCreate(d *schema.ResourceData, meta interface{}) error {
  config := meta.(*GoDaddyConfig)

  //create our id as <name>.<domain>::<type>
  id := d.Get("name").(string) + "." + d.Get("domain").(string) + "::" + d.Get("type").(string)
  d.SetId(id)

  //get the current GoDaddy state
  state, err := fetchState(
    config.key,
    config.secret,
    d.Get("domain").(string),
  )

  if err != nil {
    log.Printf("Error: %s\n", err.Error())    
    return err
  }

  //create the new DnsRecord
  record := DnsRecord {
    Data: d.Get("data").(string),
    Name: d.Get("name").(string),
    Priority: d.Get("priority").(int),
    Ttl: d.Get("ttl").(int),
    Record_type: d.Get("type").(string),
  }

  //add it to our list
  state = appendRecord(record, state)

  //write it back
  _, err = putState(
    config.key,
    config.secret,
    d.Get("domain").(string),
    state,
  )

  if err != nil {
    log.Printf("Error: %s\n", err.Error())    
    return err
  }

  return nil
}

func resourceRecordRead(d *schema.ResourceData, m interface{}) error {
  //Note: Does this function actually have a purpose?  It only seems to read from the state file, and returning
  // nil seems to be OK....
  return nil
}

func resourceRecordDelete(d *schema.ResourceData, meta interface{}) error {
  config := meta.(*GoDaddyConfig)

  //create our id as <name>.<domain>::<type>
  id := d.Get("name").(string) + "." + d.Get("domain").(string) + "::" + d.Get("type").(string)
  d.SetId(id)

  //get the current GoDaddy state
  state, err := fetchState(
    config.key,
    config.secret,
    d.Get("domain").(string),
  )

  if err != nil {
    log.Printf("Error: %s\n", err.Error())    
    return err
  }

  //create the new DnsRecord
  record := DnsRecord {
    Data: d.Get("data").(string),
    Name: d.Get("name").(string),
    Priority: d.Get("priority").(int),
    Ttl: d.Get("ttl").(int),
    Record_type: d.Get("type").(string),
  }

  //add it to our list
  state = removeRecord(record, state)

  //write it back
  _, err = putState(
    config.key,
    config.secret,
    d.Get("domain").(string),
    state,
  )

  if err != nil {
    log.Printf("Error: %s\n", err.Error())    
    return err
  }

  return nil
}
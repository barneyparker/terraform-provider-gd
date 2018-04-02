# Terraform-Provider-GD

A simple Terraform provider for GoDaddy

Although another provider already exists, I had some trouble making it work.  Unfortunately I have no experience of Go, so altering it seemed quite hard!

Instead, I wrote my own from scratch!  

It's not good (code-wise) and I am sure there are a million ways to do things a lot better with this, but it does work

## Installation:
  
### From source
To use, create a symbolic link to `~/.terraform/plugins/terraform-provider-gd`

then add the following to your ~/.terraformrc file (create if necessary)
```
gd = "$HOME/.terraform/plugins/terraform-provider-gd"
```

### Install from GitHub

```
bash <(curl -s https://raw.githubusercontent.com/barneyparker/terraform-provider-gd/master/install.sh)
```

## GoDaddy API

More information about the GoDaddy API & creation of API keys can be found at:
[GoDaddy API](https://developer.godaddy.com/)

## Declaring the Provider

In terraform, declare the provider:
```
  provider "gd" {
    key = "<API Key>"
    secret = "<Secret Key>"
  }
```

Alternatively, to keep keys out of your code, declare the provider:
```
  provider "gd" {}
```

And set environment variables:
```
  export GODADDY_API_KEY=<API Key>
  export GODADDY_API_SECRET=<Secret Key>
```

## Creating a DNS Resource

To create a resource, create a `gd_record`:
```
resource "gd_record" "test_record" {
  domain = "example.com"
  type = "A"
  name = "terraform"
  data = "8.8.8.8"
  priority = 0
  ttl = 600
}
```

which will create an `A` record for `terraform.example.com` pointing to `8.8.8.8`

Available record types are:

  * `A`
  * `AAAA`
  * `CNAME`
  * `MX`
  * `NS`
  * `TXT`

`SRV` don't work right now, and `SOA` aren't allowed by the GoDaddy API...

## State of the Project

I have no idea what I am doing!  The structure of the provider doesn't match the official providers - although I did initially start by following the Terraform guides...

The project was written in an afternoon, and while I have tried to keep the code relatively clean I was mainly focusing on getting it working rather than a perfect implementation.

There are no tests.  I found even getting logs out of Terraform to see what was going on was painful, but it really could do with some tests...

The GoDaddy API can be useless and return incorrect results occasionally.  From what I could tell, it's their API at fault.  Not a great situation, but hey....

## Collaboration

If you can improve any of the code I would welcome any pull requests.  Please do your best to explain the changes - as I mentioned above my Go experience is non-existent so I will be easily confused.

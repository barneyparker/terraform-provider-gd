provider "gd" {
  key = "<API Key>"
  secret = "<Secret Key>"
}

resource "gd_record" "test_record" {
  domain = "example.com"
  type = "A"
  name = "terraform"
  data = "8.8.8.8"
  ttl = 600
}

output "test_id" {
  value = "${gd_record.test_record.id}"
}

#!/bin/sh

tag=0.0.1
os=$(uname -s | sed -e 's/\(.*\)/\L\1/')
arch=$(uname -m)

case "$os" in
  "linux"|"darwin")
    echo "fetching latest revision..."

    if [ "$arch" == "x86_64" ]; then
      arch="amd64"
    else
      arch="386"
    fi

    filename="terraform-provider-gd.$os.$arch"

    cd /tmp && curl -fOsSL https://github.com/barneyparker/terraform-provider-gd/releases/download/v$tag/$filename.tgz 2>/dev/null

    if [ ! -f /tmp/$filename.tgz ]; then
      echo "Install failed"
      exit 1
    else 
      cd /tmp && gzip -dc $filename.tgz | tar xf -
    fi

    if [ ! -f terraform-provider-gd ]; then
      echo "Archive extract failed"
      exit 1
    fi
    ;;
  *)
    echo "Unable to download and install for this platform ($os / $arch)"
    exit 1
    ;;
esac

printf "Installing terraform-provider-gd binary..."
mkdir -p ~/.terraform/plugins && mv terraform-provider-gd ~/.terraform/plugins
if [ $? -eq 0 ]; then
  echo "Complete"
else
  echo "Failed"
  exit 1
fi

printf "Configuring Terraform..."
if [ -f ~/.terraformrc ]; then
  if [[ $(grep "gd" ~/.terraformrc) ]]; then
    echo "Complete"
  else
    echo "\nPlease add the following line to your terraform config:"
    echo ""
    echo "\tgd = \"$HOME/.terraform/plugins/terraform-provider-gd\""
    echo ""
  fi
else
  cat > ~/.terraformrc <<EOF
providers {
  gd = "$HOME/.terraform/plugins/terraform-provider-gd"
}
EOF
  echo "Complete"
fi
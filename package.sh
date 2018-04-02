#!/bin/bash

  env GOOS=$1 GOARCH=$2 go build -o terraform-provider-gd
  tar -cvzf terraform-provider-gd.$1.$2.tgz ./terraform-provider-gd
  rm ./terraform-provider-gd
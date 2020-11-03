#!/bin/sh -e

# Copied w/ love from the excellent hypnoglow/helm-s3

echo "Installing"

echo $url

mkdir -p "bin"
mkdir -p "releases/v${version}"

mv "releases/v${version}/bin/helmpush" "bin/helmpush"

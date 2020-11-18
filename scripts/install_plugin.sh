#!/bin/sh -e

version="$(cat plugin.yaml | grep "version" | cut -d '"' -f 2)"
echo "Installing helm-push ${version} ..."

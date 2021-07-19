#!/bin/bash

rm -f go.mod
rm -f go.sum
rm -rf .git
rm -f init.sh
echo "Example GO Module: gitlab.com/indev-moph/fiber-api"
read -p 'Input Go Module: ' gopack
go mod init ${gopack}

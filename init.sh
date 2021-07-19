#!/bin/bash
rm -f go.mod
rm -f go.sum
rm -rf .git
rm -f init.sh
echo "example package: gitlab.com/indev-moph/fiber-api"
read -p 'Input Go Package: ' gopack
go mod init ${gopack}

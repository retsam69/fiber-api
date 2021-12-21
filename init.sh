#!/bin/bash

rm -f go.mod
rm -f go.sum
rm -rf .git
echo "Example GO Module: gitlab.com/indev-moph/fiber-api"
read -p 'Input Go Module: ' gopack
go mod init ${gopack}
make mod-up
go install ./cmd/gmf
rm -f init.sh

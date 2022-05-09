#!/bin/bash

rm -f go.mod \
&& rm -f go.sum \
&& rm -rf .git \
&& git init \
&& echo "Example GO Module: gitlab.com/indev-moph/fiber-api" \
&& read -p 'Input Go Module: ' gopack \
&& go mod init ${gopack} \
&& make mod-up \
&& go install github.com/attapon-th/gmf/cmd/gmf@latest \
&& rm -f init.sh

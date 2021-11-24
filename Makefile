.PHONY: all
BINARY=AppMain
VERSION=0.0.1
BUILD=faa91d431795e148ca10c3f6613000921a57dec2

# go main file
GOMAINFILE=main.go
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"
GO_MODULE=`cat go.mod | grep -m1 module | sed 's/^module \(.*\)$$/\1/'`
GIT_REGISTRY_URL=registry.${GO_MODULE}

SSH_NAME=ssh_name
SERVICE_NAME=docker_service_name

dev:
	go run ${LDFLAGS} main.go -c dev.yaml

swagger: 
	rm -rf ./docs
	swag init --parseInternal --generatedTime
	make swag2openapi

swagger-prd:
	rm -rf ./docs
	swag init --parseInternal --generatedTime -g main.go-prd
	make swag2openapi

swag2openapi:
	rm -rf ./docs/docs.go
	curl -X POST "https://converter.swagger.io/api/convert" \
	-H "accept: application/json" \
	-H "Content-Type: application/json" \
	-d @./docs/swagger.json \
	> docs/openapi-${VERSION}.json


mod-up:
	go mod tidy
	@echo ">> updating Go dependencies"
	@for m in $$(go list -mod=readonly -m -f '{{ if and (not .Indirect) (not .Main)}}{{.Path}}{{end}}' all); do \
		go get $$m; \
	done
	go mod vendor

mod:
	go mod tidy
	go mod vendor

docker-build:
	gmf ibv
	-git add .
	-git cpush
	make swagger-prd
	docker build \
	-t ${GIT_REGISTRY_URL}:latest .

docker-addtags:
	docker tag ${GIT_REGISTRY_URL}:latest ${GIT_REGISTRY_URL}:${VERSION} 

docker-push:
	docker push ${GIT_REGISTRY_URL}:latest 

build-in-docker:
	CGO_ENABLED=0 GOOS=linux go build \
	-a -installsuffix cgo ${LDFLAGS} \
	-o ${BINARY} ${GOMAINFILE}
	
	
move-in-docker:
	mv ${BINARY} /app/${BINARY} 
	mkdir -p /app/docs
	mv docs/openapi-${VERSION}.json /app/docs/openapi-${VERSION}.json
	mv default.yml /app/default.yaml


server-up:
	@echo "Server Up ${SSH_NAME}"
	ssh ${SSH_NAME} "docker pull ${GIT_REGISTRY_URL}:latest; \
	docker service update --force ${SERVICE_NAME};"

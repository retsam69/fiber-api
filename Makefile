.PHONY: all
BINARY=AppMain
VERSION=0.0.0
BUILD=ff8f8ac4f4a3e43945b76621c13fc9e99878f1ec





# go main file
GOMAINFILE=main.go
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"
GO_MODULE=`cat go.mod | grep -m1 module | sed 's/^module \(.*\)$$/\1/'`
GIT_REGISTRY_URL="registry.${GO_MODULE}"

SSH_NAME=ssh_name
SERVICE_NAME=docker_service_name

dev:
	go run ${LDFLAGS} main.go -c dev.yaml

swagger:
	rm -rf ./docs
	swag init --parseInternal --generatedTime
	rm -rf ./docs/docs.go
	curl -X POST "https://converter.swagger.io/api/convert" \
	-H "accept: application/json" \
	-H "Content-Type: application/json" \
	-d @./docs/swagger.json \
	> docs/openapi-${VERSION}.json


update-go-deps:
	go mod tidy
	@echo ">> updating Go dependencies"
	@for m in $$(go list -mod=readonly -m -f '{{ if and (not .Indirect) (not .Main)}}{{.Path}}{{end}}' all); do \
		go get $$m; \
	done

mod: update-go-deps	
	go mod tidy
	go mod vendor

git:
	-git add . 
	-git commit -m "update" 
	git push

version-up:
	gmf i
	gmf b
	gmf v

docker-build: version-up git swagger
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


server-up:
	@echo "Server Up ${SSH_NAME}"
	ssh ${SSH_NAME} "docker pull ${GIT_REGISTRY_URL}:latest; \
	docker service update --force ${SERVICE_NAME};"

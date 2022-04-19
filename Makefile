BINARY=AppMain
VERSION=0.0.0
BUILD=b7f8cd0e388bed7781c313eaf58034e2ba911237

# go main file
GOMAINFILE=main.go
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"
GO_MODULE=`cat go.mod | grep -m1 module | sed 's/^module \(.*\)$$/\1/'`
GIT_REGISTRY_URL=registry.${GO_MODULE}
SSH_NAME=ssh_config_name
REMOTE_PATH=/home/attapon/${BINARY}

dev:
	go run ${LDFLAGS} main.go -c dev.env

swagger: 
	swag fmt
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
	@for m in $$(go list -mod=readonly -m -f '{{ if and (not .Indirect) (not .Main)}}{{.Path}}{{end}}' all); do go get $$m; done
	go mod vendor

mod:
	go mod tidy
	go mod vendor

docker-build:
	gmf ibv
	-git add .
	-git cpush
	make swagger-prd
	echo "Start Build docker image: ${GIT_REGISTRY_URL}"
	docker build \
	--build-arg ENTRYPOINTNAME=${BINARY} \
	-t ${GIT_REGISTRY_URL}:latest .

docker-addtags:
	docker tag ${GIT_REGISTRY_URL}:latest ${GIT_REGISTRY_URL}:${VERSION} 

docker-push:
	docker push ${GIT_REGISTRY_URL}:latest 

docker-clean:
	# remove docker images <none>:<none>
	-docker rmi $(docker images -f "dangling=true" -q)

build-in-docker:
	CGO_ENABLED=0 GOOS=linux go build \
	-a -installsuffix cgo ${LDFLAGS} \
	-o ${BINARY} ${GOMAINFILE}
	
move-in-docker:
	mv ${BINARY} /app/${BINARY} 
	mkdir -p /app/docs
	mv docs/openapi-${VERSION}.json /app/docs/openapi-${VERSION}.json



server-sync:
	ssh ${SSH_NAME} "mkdir -p ${REMOTE_PATH}/logs"
	scp deployments/docker-compose.yml ${SSH_NAME}:${REMOTE_PATH}/docker-compose.yml

server-up:
	@echo "Server Up ${SSH_NAME}"
	ssh ${SSH_NAME} "cd ${REMOTE_PATH}; \
	docker-compose pull; \
	docker-compose up -d && docker-compose logs -f"

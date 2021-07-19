.PHONY: all
ifndef BINARY
BINARY=AppMain
endif
ifndef VERSION
VERSION=0.0.1
endif
ifndef BUILD
BUILD=`git rev-parse HEAD`
endif
# go main file
GOMAINFILE=main.go
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"
GO_MODULE=`cat go.mod | grep -m1 "module " | sed 's/^module \(.*\)$/\1/'`
GIT_REGISTRY_URL="registry.${GO_MODULE}"


dev:
	USER_ADMIN=1234
	APP_DEV=true \
	APP_LISTEN=127.0.0.1:8888 \
	APP_PREFIX=/api \
	go run ${LDFLAGS} main.go

swagger:
	swag init

update-go-deps:
	@echo ">> updating Go dependencies"
	@for m in $$(go list -mod=readonly -m -f '{{ if and (not .Indirect) (not .Main)}}{{.Path}}{{end}}' all); do \
		go get $$m; \
	done

mod: update-go-deps	
	go mod tidy
	go mod vendor

git-cp:
	git add .
	git commit --amend
	git push

docker-build:
	docker build \
	--build-arg BUILDDOCKER=$(BUILD) \
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
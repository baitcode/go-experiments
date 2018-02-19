.PHONY: build

define linux_build
	mkdir -p build/linux
	@GOOS=linux go build -o build/linux/$2 -i $1
endef

define osx_build
	mkdir -p build/osx
	@GOOS=darwin go build -o build/osx/$2 -i $1
endef

build:
	$(call osx_build,./main.go,management)

sls-deploy:
	@cd ./deployments/serverless && ./deploy.sh

sls-remove:
	@cd ./deployments/serverless && ./remove.sh

test:
	@TEST_RUN=true PROJECT_ROOT=$(shell pwd) go test ./... | grep -v "no test files"
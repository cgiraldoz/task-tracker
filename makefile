APP=task-tracker
REGISTRY=ghcr.io
MAIN_GO=main.go
TAG=latest
USER=cgiraldoz

ifeq ($(OS), Windows_NT)
    EXECUTABLE=$(APP).exe
    COMMIT_SHA := $(shell powershell -Command "git rev-parse --short HEAD")
    DOCKER_LOGIN := docker login $(REGISTRY) -u $(USER) -p $(GH_REGISTRY_TOKEN)
else
    EXECUTABLE=$(APP)
    COMMIT_SHA := $(shell git rev-parse --short HEAD)
    DOCKER_LOGIN := echo $(GH_REGISTRY_TOKEN) | docker login $(REGISTRY) -u $(USER) --password-stdin
endif

.PHONY: build run clean help

## build: Build the application
build: clean
	go build -o $(EXECUTABLE) $(MAIN_GO)

## run: Run the application
run:
	go run $(MAIN_GO)

## clean: Clean the application
clean:
	go clean

docker-login:
	@$(DOCKER_LOGIN)

## docker-build: Build the Docker image
docker-build: build
	docker build -t $(REGISTRY)/$(USER)/$(APP):$(TAG) .

## docker-push: Push the Docker image
docker-push: docker-build
	docker tag $(REGISTRY)/$(USER)/$(APP):$(TAG) $(REGISTRY)/$(USER)/$(APP):$(COMMIT_SHA)
	docker push $(REGISTRY)/${USER}/$(APP):$(TAG)
	docker push $(REGISTRY)/$(USER)/$(APP):$(COMMIT_SHA)

## docker-run: Run the Docker image
docker-run: docker-build
	docker run -it --rm $(REGISTRY)/$(USER)/$(APP):$(TAG)

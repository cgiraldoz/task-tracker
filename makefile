APP=task-tracker
COMMIT_SHA := $(shell powershell -Command "git rev-parse --short HEAD")


build:
	@echo $(COMMIT_SHA)
	go build -o $(APP).exe main.go

run:
	go run main.go

clean:
	go clean

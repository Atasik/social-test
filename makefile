.PHONY:
.SILENT:

build:
    eval $(minikube docker-env) && docker build -t social:latest .

run: build
    docker compose up social

test:
    go test -v ./...
run:
	go run main.go

build:
	go mod tidy && go build -o bin/

docker-build-run:
	docker build -t go-gin-starter:0.1 .;
	docker run --name=go-gin-starter --rm -p 8080:8080 go-gin-starter:0.1 

docker-run:
	docker run --name=go-gin-starter --rm -p 8080:8080 go-gin-starter:0.1 
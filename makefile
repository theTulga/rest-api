dev:
	go run ./src/main.go
build:
	go build -o="./bin/main" ./src/main.go
run-prod:
	make build
	./main -prod=true
run: 
	go run ./src/main.go
	./main -prod=true
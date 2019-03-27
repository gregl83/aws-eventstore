.PHONY: build clean deploy

build:
	dep ensure -v
	env GOOS=linux go build -ldflags="-s -w" -o bin/migrate migrate/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/write write/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/read read/main.go
	sls sam export -o template.yml

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose

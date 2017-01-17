all: format test
	go run main.go

test:
	cd factorial && go test
	go test

format:
	find . -name \*.go -exec go fmt '{}' \;

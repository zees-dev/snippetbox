all: clean test build serve

clean:
	rm -rf server

build:
	go build -o server cmd/web/main.go

serve:
	go run cmd/web/main.go

test:
	echo "todo..."
.PHONY: test

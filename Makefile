all: clean test build serve

clean:
	rm -rf server

build:
	go build -o server main.go

serve:
	go run main.go

test:
	echo "todo..."
.PHONY: test

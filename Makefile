all: build

BINARY=main

build:
	go build -o $(BINARY)

run: 
	-@go build -o $(BINARY)
	-@./$(BINARY)

clean:
	-@rm -f $(BINARY)

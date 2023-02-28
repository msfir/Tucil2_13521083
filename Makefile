all: build

build:
	@cd src && go build
	@mkdir -p bin
	@mv src/pairit bin

clean:
	@rm -rf bin

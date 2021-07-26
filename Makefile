VERSION = "0.0.1"

build:
	go build -v -o bin/commando main/commando.go

install:
	go install main/commando.go

clean:
	@rm -rf bin/commando
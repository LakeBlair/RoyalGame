binaries = main

build:
	go build main.go

run:
	go run main.go

clean:
	rm -f $(binaries)
target = ./bin/onion-node

all:
	mkdir -p bin
	go build src/*.go
	mv onion-node bin

clean:
	rm -f $(target)

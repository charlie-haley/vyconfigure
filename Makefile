test:
	go test -v ./...

lint:
	golangci-lint run

dist:
	mkdir dist

clean:
	rm -Rf dist

build: clean dist
	GO111MODULE=on go build -v  -ldflags="-s -w" -trimpath -o ./dist/vyconfigure

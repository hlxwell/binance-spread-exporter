# run
run:
	go run -v ./...

# build
build:
	mkdir build
	go build -o build/spread_exporter -v ./...

# test
test:
	go test -v ./...

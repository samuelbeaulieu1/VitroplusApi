run:
	go run main.go

build:
	rm -rf build/
	go build -o build/
	cp config.json build/

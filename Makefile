manifest:
	@go run github.com/akavel/rsrc@v0.10.2 -manifest setup.exe.manifest -o setup.exe.syso

build: manifest
	@mkdir -p build
	@GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o build/setup.exe .

run: build
	@./build/setup.exe

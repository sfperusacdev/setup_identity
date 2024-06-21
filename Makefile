manifest:
	@rsrc.exe -manifest setup.exe.manifest -o setup.exe.syso

build: manifest
	@go build -o setup.exe .

run:
	@go build -o setup.exe .
	@setup.exe
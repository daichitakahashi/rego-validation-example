test:
	opa test -v policy

build:
	go run ./cmd/build -src policy -wasm artifact/policy.wasm -bundle go/bundle.tar.gz
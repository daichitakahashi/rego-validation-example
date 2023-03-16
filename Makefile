test:
	opa test -v policy

build:
	mkdir -p artifact
	opa build -t wasm -b policy -e validation/email -e validation/domain -o artifact/bundle.tar.gz
	tar zfx artifact/bundle.tar.gz -C artifact /policy.wasm
	rm artifact/bundle.tar.gz
	opa build -t rego -b policy -e validation/email -e validation/domain -o go/bundle.tar.gz --optimize=2

build-js: export GOOS=js
build-js: export GOARCH=wasm
build-js: wasm_exec.js
	go build -o ebitentest.wasm ./cmd/main.go

wasm_exec.js:
	cp ${GOROOT}/misc/wasm/wasm_exec.js .

run-local: build-js
	go run cmd/testserver/main.go

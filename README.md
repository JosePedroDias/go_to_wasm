# go to wasm

https://go.dev/wiki/WebAssembly
https://pkg.go.dev/syscall/js

```
GOOS=js GOARCH=wasm go build -o main.wasm
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
```

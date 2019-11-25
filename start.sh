if [ ! -f 'wasm_exec.js' ]; then
    cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
fi

GOOS=js GOARCH=wasm go build -o main.wasm
goexec 'http.ListenAndServe(`:8888`, http.FileServer(http.Dir(`.`)))'

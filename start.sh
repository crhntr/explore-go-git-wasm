main() {
  trap "rm -f main.wasm wasm_exec.js" EXIT

  if [ ! -f 'wasm_exec.js' ]; then
      cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
  fi

  if [ ! -f 'main.wasm' ]; then
      GOOS=js GOARCH=wasm go build -o main.wasm
  fi

  sleep 2 && open 'http://localhost:8888/' &
  goexec 'http.ListenAndServe(`:8888`, http.FileServer(http.Dir(`.`)))'
}

main

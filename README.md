# servehere
Simple http static file server that can be used to serve the current working directory as the web root. 
It's can be used similarly to python 2's `SimpleHTTPServer` or python 3's `http.server`.

## Install
`go install github.com/quillaja/servehere`

## Run
Assuming the `$GOPATH/bin` is in your `$PATH`, run:
`$ servehere`

Optional parameters are:

| Flag | Description | Default |
|---|---|---|
| p | port to listen on | 8000 |
| a | bind address | localhost |
| v | verbose output | true |

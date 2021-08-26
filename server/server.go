package main

import (
	"io"
	"net/http"
	"net/rpc"
	"rpc-demo/common"
)

func main() {

	// create a `*College` object
	mit := common.ReadFile()

	// register `mit` object with `rpc-demo.DefaultServer`
	rpc.Register(mit)

	// register an HTTP handler for RPC communication on `http.DefaultServeMux` (default)
	// registers a handler on the `rpc-demo.DefaultRPCPath` endpoint to respond to RPC messages
	// registers a handler on the `rpc-demo.DefaultDebugPath` endpoint for debugging
	rpc.HandleHTTP()

	// sample test endpoint
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, "RPC SERVER LIVE!")
	})

	// listen and serve default HTTP server
	http.ListenAndServe(":9000", nil)

}

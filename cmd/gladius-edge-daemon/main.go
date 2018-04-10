package main

import (
	"fmt"
	"gladius-edge-daemon/init/manager"
	"net"
	"net/http"
	"net/rpc"

	"github.com/powerman/rpc-codec/jsonrpc2"
	"github.com/valyala/fasthttp"
)

// Create transport types
type RPCData struct {
}

type HTTPData struct {
}

type GladiusEdge struct {
	rpcData RPCData
}

// Start - Start the gladius edge node
func (*GladiusEdge) Start(vals [2]int, res *string) error {
	*res = "Not Implemented"
	return nil
}

// Stop - Stop the gladius edge node
func (*GladiusEdge) Stop(vals [2]int, res *string) error {
	*res = "Not Implemented"
	return nil
}

// Status - Get the current status of the network node
func (*GladiusEdge) Status(vals [2]int, res *string) error {
	*res = "Not Implemented"
	return nil
}

// Main entry-point for the service
func main() {
	// Define some variables
	name, displayName, description :=
		"GladiusEdgeDaemon",
		"Gladius Network (Edge) Daemon",
		"Gladius Network (Edge) Daemon"

	// Run the function "run" as a service
	manager.RunService(name, displayName, description, run)
}

// Start a web server
func run() {
	fmt.Println("Starting...")
	// Create some strucs so we can pass info between goroutines
	rpcData := &RPCData{}
	// httpData := &HTTPData{}

	// Create some channels to communicate out of the threads
	contentChannel, rpcChannel := make(chan string), make(chan string)

	// Create a content server goroutine
	go fasthttp.ListenAndServe(":8080", requestHandler(contentChannel, rpcData))

	// Register RPC methods
	rpc.Register(&GladiusEdge{})

	// Setup HTTP handling for RPC on port 5000
	http.Handle("/rpc", jsonrpc2.HTTPHandler(nil))
	lnHTTP, err := net.Listen("tcp", ":5000")
	if err != nil {
		panic(err)
	}
	defer lnHTTP.Close()
	go http.Serve(lnHTTP, nil)

	fmt.Println("Started RPC server and HTTP server.")

	// Forever check through the channels on the main thread
	for {
		select {
		case i := <-contentChannel: // If it can be assigned to a variable
			fmt.Printf("it's a %q", i)
		case i := <-rpcChannel: // If it can be assigned to a variable
			fmt.Printf("it's a %q", i)
		}
	}
}

// Return a function like the one fasthttp is expecting
func requestHandler(contentChannel chan string, rpcData *RPCData) func(ctx *fasthttp.RequestCtx) {
	// The actual serving function
	return func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/content":
			contentHandler(ctx)
			contentChannel <- "test" // Write to the channel (will likely be an int)
		default:
			ctx.Error("Unsupported path", fasthttp.StatusNotFound)
		}
	}
}

func contentHandler(ctx *fasthttp.RequestCtx) {
	// URL format like /content?website=REQUESTED_SITE
	website := string(ctx.QueryArgs().Peek("website"))

	// // TODO: Make this serve the appropriate JSON
	fmt.Fprintf(ctx, "Hi there! You asked for %q", website)
}

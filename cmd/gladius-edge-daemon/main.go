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

// RPCOut - Transport struct
type RPCOut struct {
	httpState chan bool
}

// HTTPOut - Transport struct
type HTTPOut struct {
}

// GladiusEdge - Entry for the RPC interface. Methods take the form GladiusEdge.Method
type GladiusEdge struct {
	rpcOut *RPCOut
}

// Start - Start the gladius edge node
func (g *GladiusEdge) Start(vals [2]int, res *string) error {
	g.rpcOut.httpState <- true
	*res = "Started the server"
	return nil
}

// Stop - Stop the gladius edge node
func (g *GladiusEdge) Stop(vals [2]int, res *string) error {
	g.rpcOut.httpState <- false
	*res = "Stopped the server"
	return nil
}

// Status - Get the current status of the network node
func (g *GladiusEdge) Status(vals [2]int, res *string) error {
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
	rpcOut := &RPCOut{make(chan bool)}
	httpOut := &HTTPOut{}

	// Content server stuff below

	// Listen on 8080
	lnContent, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	// Create a content server
	server := fasthttp.Server{Handler: requestHandler(httpOut)}
	// Serve the content
	defer lnContent.Close()
	go server.Serve(lnContent)

	// RPC Stuff below

	// Register RPC methods
	rpc.Register(&GladiusEdge{rpcOut: rpcOut})
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
		case state := <-(*rpcOut).httpState: // If it can be assigned to a variable
			if state {
				lnContent, err := net.Listen("tcp", ":8080")
				if err != nil {
					panic(err)
				}
				go server.Serve(lnContent)
				fmt.Println("Started")
			} else {
				lnContent.Close()
				fmt.Println("Stopped")
			}
		}
	}
}

// Return a function like the one fasthttp is expecting
func requestHandler(httpOut *HTTPOut) func(ctx *fasthttp.RequestCtx) {
	// The actual serving function
	return func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/content":
			contentHandler(ctx)
			// TODO: Write stuff to pass back to httpOut
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

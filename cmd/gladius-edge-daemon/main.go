package main

import (
	"fmt"
	"gladius-edge-daemon/init/manager"

	"github.com/valyala/fasthttp"
)

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
	// pass plain function to fasthttp
	fasthttp.ListenAndServe(":8080", requestHandler)
}

// the corresponding fasthttp request handler
func requestHandler(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "/content":
		contentHandler(ctx)
	default:
		ctx.Error("Unsupported path", fasthttp.StatusNotFound)
	}
}

func contentHandler(ctx *fasthttp.RequestCtx) {
	// URL format like /content?website=REQUESTED_SITE
	website := string(ctx.QueryArgs().Peek("website"))

	// // TODO: Make this serve the appropriate JSON
	fmt.Fprintf(ctx, "Hi there! You asked for %q", website)
}

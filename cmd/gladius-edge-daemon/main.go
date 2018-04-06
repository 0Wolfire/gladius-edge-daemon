package main

import (
	"fmt"
	"gladius-edge-daemon/init/manager"
	"net/http"

	"github.com/osamingo/jsonrpc"
	"github.com/valyala/fasthttp"
)

var contentChannel, rpcChannel = make(chan string), make(chan string)

type (
	HandleParamsResulter interface {
		jsonrpc.Handler
		Params() interface{}
		Result() interface{}
	}
	UserService struct {
		SignUpHandler HandleParamsResulter
		LogInHandler  HandleParamsResulter
	}
)

func NewUserService() *UserService {
	return &UserService{
		// Initialize handlers
	}
}

// Create transport types
type RPCData struct {
}

type HTTPData struct {
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

	// Create a content server goroutine
	go fasthttp.ListenAndServe(":8080", requestHandler(contentChannel, rpcData))

	mr := jsonrpc.NewMethodRepository()
	us := NewUserService()

	mr.RegisterMethod("UserService.SignUp", us.SignUpHandler, us.SignUpHandler.Params(), us.SignUpHandler.Result())
	mr.RegisterMethod("UserService.LogIn", us.LoginHandler, us.LoginHandler.Params(), us.LoginHandler.Result())

	go http.ListenAndServe(":5000", http.DefaultServeMux)

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

func rpcHandler(rpcChannel chan string, httpData *HTTPData) {

}

func contentHandler(ctx *fasthttp.RequestCtx) {
	// URL format like /content?website=REQUESTED_SITE
	website := string(ctx.QueryArgs().Peek("website"))

	// // TODO: Make this serve the appropriate JSON
	fmt.Fprintf(ctx, "Hi there! You asked for %q", website)
}

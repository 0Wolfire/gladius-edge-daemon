var express = require('express');
var jayson = require('jayson');
var bodyParser = require('body-parser');

// Create an express app for the RPC Server
var rpcApp = express();

var staticApp; // Placeholder for the static app server
var staticServer; // Store the server object from app.listen
var running = false; // Running state of the static content app

// Set up parsing
rpcApp.use(bodyParser.urlencoded({
  extended: true
}));
rpcApp.use(bodyParser.json());

// Build Jayson (JSON-RPC server)
var rpcServer = jayson.server({
  start: function(args, callback) {
    staticServer = staticApp.listen(8080); // Start the app
    running = true;
    callback(null, "Started server");
  },
  stop: function(args, callback) {
    if (staticServer) { // Make sure the server is started
      staticServer.close(); // Stop the app
      running = false;
      callback(null, "Stopped server");
    } else {
      callback(null, "Server was not running so can't be stopped.");
    }
  },
  status: function(args, callback) {
    callback(null, {
      running: running // Return the current running status
    })
  },
  addContent: function(args, callback){
    // TODO: Add new content to the "__dirname + '/cdn_content/'" folder
    callback(null, "Added new content")
  }
});

// Add the RPC endpoint
rpcApp.post('/rpc', rpcServer.middleware());

// Export a start function
exports.start = function(app) {
  staticApp = app;
  rpcApp.listen(5000);
}

var express = require('express');
var jayson = require('jayson');
var bodyParser = require('body-parser');
var fs = require('fs');

// Create an express app for the RPC Server
var rpcApp = express();

var staticApp; // Placeholder for the static app server
var staticServer; // Store the server object from app.listen
var running = false; // Running state of the static content app
var cdnData = {}; // Stores {filename: base64 encoded file}

// Check the cdn_content directory for new bundle files
function checkForBundles() {
  fs.readdir(__dirname + "/../cdn_content/", (err, files) => {
    files.forEach(file => {
      if (file.endsWith(".json"))
        cdnData[file] = JSON.parse(fs.readFileSync(__dirname +
          "/../cdn_content/" + file));
    });
  });
}

// Set up parsing
rpcApp.use(bodyParser.urlencoded({
  extended: true
}));
rpcApp.use(bodyParser.json());

function reloadData() {
  var directory = __dirname + '/cdn_content/';
}

// Build Jayson (JSON-RPC server)
var rpcServer = jayson.server({
  start: function(args, callback) {
    if (!running) {
      staticServer = staticApp.listen(8080); // Start the app
      running = true;
      callback(null, "Started server");
    } else {
      callback(null, "Server already running");
    }

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
      //TODO: Talk about configuring a hostname for SSL in future
      running: running // Return the current running status
    })
  },
  reloadContent: function(args, callback) {
    checkForBundles();
    callback(null, "Checked for new bundles")
  }
});

// Load up the bundles at start
checkForBundles();

// Add the RPC endpoint
rpcApp.post('/rpc', rpcServer.middleware());

// Export a start function
exports.start = function(app) {
  staticApp = app;
  rpcApp.listen(5000);
}

exports.getCDNData = function(bundleName) {
  return cdnData[bundleName];
}

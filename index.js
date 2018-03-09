var express = require('express');
var rpc = require('./rpc/rpc.js')

var staticApp = express(); // Express app to serve static content

rpc.start(staticApp);

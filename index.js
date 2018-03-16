var express = require('express');
var cors = require('cors');
var rpcController = require('./controller/controller.js')

var staticApp = express(); // Express app to serve static content
staticApp.use(cors());

staticApp.get('/content_bundle/', function(req, res) {
  return res.json(rpcController.getCDNData());
});

rpcController.start(staticApp);

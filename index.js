var express = require('express');
var rpcController = require('./rpc/rpc.js')

var staticApp = express(); // Express app to serve static content

staticApp.get('/content/:name', function(req, res) {
  res.sendFile(req.params.name, {
    root: __dirname + '/cdn_content/'
  });
});

rpcController.start(staticApp);

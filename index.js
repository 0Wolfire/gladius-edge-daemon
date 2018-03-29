#!/usr/bin/env node

var express = require('express');
var cors = require('cors');
var rpcController = require('./controller/controller.js');
var pjson = require('./package.json');

var staticApp = express(); // Express app to serve static content
staticApp.use(cors());

staticApp.get('/content_bundle/', function(req, res) {
  return res.json(rpcController.getCDNData("bundle.json"));
});

staticApp.get('/status/', function(req, res){
  return res.json({version: pjson.version, status: "up"});
});

rpcController.start(staticApp);

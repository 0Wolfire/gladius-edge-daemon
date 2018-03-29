# Gladius Edge Node Daemon

The edge node daemon to be installed alongside the control daemon.

### Install
Clone this repo and run `npm install`

### Run
Run `node index.js` in the project directory

### Test the RPC server

```
$ HDR='Content-type: application/json'

$ MSG='{"jsonrpc": "2.0", "method": "start", "id": 1}'
$ curl -H $HDR -d $MSG http://localhost:5000/rpc
{"jsonrpc":"2.0","result":"Started server","id":1}

$ MSG='{"jsonrpc": "2.0", "method": "stop", "id": 1}'
$ curl -H $HDR -d $MSG http://localhost:5000/rpc
{"jsonrpc":"2.0","result":"Stopped server","id":1}

$ MSG='{"jsonrpc": "2.0", "method": "status", "id": 1}'
$ curl -H $HDR -d $MSG http://localhost:5000/rpc
{"jsonrpc":"2.0","result":{"running": false},"id":1}
```

### Set up content delivery

Add bundle.json to cdn_content directory (in future releases this will be done
  automatically)

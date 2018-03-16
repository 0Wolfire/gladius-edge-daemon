# Gladius Edge Node Daemon

The edge node daemon to be installed alongside the control daemon.

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

Add files to cdn_content directory

Access them at localhost:8080/content/file_name   

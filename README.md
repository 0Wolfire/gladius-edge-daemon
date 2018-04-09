# Gladius Edge Node Daemon

The edge node daemon to be installed alongside the control daemon.

### Install
Install [go](https://golang.org/doc/install)

Compile and install the main.go file in cmd/gladius-edge-daemon/ with `go isntall cmd/gladius-edge-daemon/main.go`

##### Some untested stuff with services
Install the service on your machine with: `gladius-edge-daemon install`
Start with: `gladius-edge-daemon start`
Stop with: `gladius-edge-daemon stop`


### Run
Run the executable created by the above step with `gladius-edge-daemon`
(Or use the steps above and make it a service)

### Test the RPC server (Not implemented yet)
```bash
$ HDR1='Content-type: application/json'
$ HDR2='Accept: application/json'

$ MSG='{"jsonrpc": "2.0", "method": "start", "id": 1}'
$ curl -H $HDR1 -H $HDR2 -d $MSG http://localhost:5000/rpc
{"jsonrpc":"2.0","result":"Started server","id":1}

$ MSG='{"jsonrpc": "2.0", "method": "stop", "id": 1}'
$ curl -H $HDR1 -H $HDR2 -d $MSG http://localhost:5000/rpc
{"jsonrpc":"2.0","result":"Stopped server","id":1}

$ MSG='{"jsonrpc": "2.0", "method": "status", "id": 1}'
$ curl -H $HDR1 -H $HDR2 -d $MSG http://localhost:5000/rpc
{"jsonrpc":"2.0","result":{"running": false},"id":1}
```

### Set up content delivery

Bundle files will be fetched from the masternode and loaded into RAM

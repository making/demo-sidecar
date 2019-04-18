# Sidecar demo on Cloud Foundry

Sidecar is suppoted since CAPI 1.79.0 / CC API Version 2.134.0.
Make sure `cf curl /v2/info | jq -r .api_version` is 2.134.0+.

```
GOOS=linux go build ./cmd/sidecar/
cf push --no-start
cf curl /v3/apps/$(cf app demo-sidecar --guid)/sidecars -d '{"name":"demo","command":"./sidecar","process_types": ["web"]}'
cf start demo-sidecar
```

```
$ curl https://demo-sidecar.yourcf.example.com -w '\n'
Sidecar received your data
```

```
$ curl https://demo-sidecar.yourcf.example.com/kill-sidecar
```

```
   2019-04-18T16:29:33.85+0900 [APP/PROC/WEB/SIDECAR/DEMO/0] OUT Exit status 137
   2019-04-18T16:29:33.85+0900 [CELL/SSHD/0] OUT Exit status 0
   2019-04-18T16:29:33.87+0900 [APP/PROC/WEB/0] OUT Exit status 143
```

```
$ curl https://demo-sidecar.yourcf.example.com/kill-main
```

```
   2019-04-18T16:32:22.04+0900 [APP/PROC/WEB/0] OUT Exit status 137
   2019-04-18T16:32:22.05+0900 [CELL/SSHD/0] OUT Exit status 0
   2019-04-18T16:32:22.06+0900 [APP/PROC/WEB/SIDECAR/DEMO/0] OUT Exit status 143
```


This sample app is a clone of 
* https://github.com/cloudfoundry/capi-bara-tests/tree/master/assets/sidecar
* https://github.com/cloudfoundry/capi-bara-tests/tree/master/assets/sidecar-dependent
# Sidecar demo on Cloud Foundry

Sidecar is suppoted since CAPI 1.79.0 / CC API Version 2.134.0.
Make sure `cf curl /v2/info | jq -r .api_version` is 2.134.0+.

```
GOOS=linux go build ./cmd/sidecar/
cf push --no-start
cf curl /v3/apps/$(cf app demo-sidecar --guid)/sidecars -d '{"name":"demo","command":"./sidecar","process_types": ["web", "worker"]}'
cf start demo-sidecar
```

```
$ curl https://demo-sidecar.yourcf.example.com -w '\n'
Sidecar received your data
```
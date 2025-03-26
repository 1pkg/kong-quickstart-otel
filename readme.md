## Kong quickstart opentelemetry plugin

Update the `otelconfig.yaml` to start OTel collector and few services.

```bash
docker compose up
```

Initialize Kong with opentelemetry [pluging](https://docs.konghq.com/hub/kong-inc/opentelemetry/).

```bash
./kong.sh
```

Setup the services routing in Kong.

```bash
./services.sh
```

Send a request to the services.

```bash
curl http://localhost:8000/a                                                                                           
```
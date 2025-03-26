## Kong quickstart opentelemetry plugin

Update the `otelconfig.yaml` to start OTel collector and few services.

```bash
docker compose build --no-cache && ELASTIC_APM_SERVER_URL="https://my-apm-server-url:443" ELASTIC_APM_SECRET_TOKEN="REDACTED" docker compose up -d
```

Initialize Kong with opentelemetry [plugin](https://docs.konghq.com/hub/kong-inc/opentelemetry/).

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
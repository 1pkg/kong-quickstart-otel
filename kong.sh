#!/usr/bin/env sh

curl -Ls get.konghq.com/quickstart | bash -s -- -e KONG_TRACING_INSTRUMENTATIONS=all -e KONG_TRACING_SAMPLING_RATE=1.0

curl -X POST \
  --url http://localhost:8001/plugins \
  --header 'Content-Type: application/json' \
  --data '{
	"name": "opentelemetry",
	"config": {
		"traces_endpoint": "http://172.17.0.1:4318/v1/traces",
		"propagation": {
		  "inject": ["w3c"],
		  "extract": ["w3c"],
		  "default_format": "w3c"
		},
		"resource_attributes": {
      "service.name": "kong-dev"
    }
	}
}'

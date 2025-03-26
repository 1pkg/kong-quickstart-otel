#!/usr/bin/env sh

curl -Ls get.konghq.com/quickstart | sh -s -- -t 3.1.1.1-alpine -e KONG_OPENTELEMETRY_TRACING=all -e KONG_LOG_LEVEL=debug

curl --request POST \
  --url http://localhost:8001/plugins \
  --header 'Content-Type: application/json' \
  --data '{
	"name": "opentelemetry",
	"config": {
		"endpoint": "http://kong-demo-opentelemetry-collector-1:4318/v1/traces"
	}
}'
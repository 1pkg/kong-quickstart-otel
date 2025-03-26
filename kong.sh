#!/usr/bin/env sh

curl -Ls get.konghq.com/quickstart | bash

curl --request POST \
  --url http://localhost:8001/plugins \
  --header 'Content-Type: application/json' \
  --data '{
	"name": "opentelemetry",
	"config": {
		"endpoint": "http://172.17.0.1:4318/v1/traces"
	}
}'
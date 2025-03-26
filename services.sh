#!/usr/bin/env sh

curl --request POST \
  --url http://localhost:8001/services/ \
  --data name=servicea \
  --data url=http://kong-demo-servicea-1
  
curl --request POST \
  --url http://localhost:8001/services/servicea/routes \
  --data paths=/a

curl --request POST \
  --url http://localhost:8001/services/ \
  --data name=serviceb \
  --data url=http://kong-demo-serviceb-1
  
curl --request POST \
  --url http://localhost:8001/services/serviceb/routes \
  --data paths=/b

curl --request POST \
  --url http://localhost:8001/services/ \
  --data name=servicec \
  --data url=http://kong-demo-servicec-1
  
curl --request POST \
  --url http://localhost:8001/services/servicec/routes \
  --data paths=/c
#!/usr/bin/env sh

curl --request POST \
  --url http://localhost:8001/services/ \
  --data name=servicea \
  --data url=http://host.docker.internal:10080
  
curl --request POST \
  --url http://localhost:8001/services/servicea/routes \
  --data paths=/a

curl --request POST \
  --url http://localhost:8001/services/ \
  --data name=serviceb \
  --data url=http://host.docker.internal:10081
  
curl --request POST \
  --url http://localhost:8001/services/serviceb/routes \
  --data paths=/b

curl --request POST \
  --url http://localhost:8001/services/ \
  --data name=servicec \
  --data url=http://host.docker.internal:10082
  
curl --request POST \
  --url http://localhost:8001/services/servicec/routes \
  --data paths=/c

# py

curl --request POST \
  --url http://localhost:8001/services/ \
  --data name=servicepy \
  --data url=http://host.docker.internal:10083
  
curl --request POST \
  --url http://localhost:8001/services/servicepy/routes \
  --data paths=/py

curl --request POST \
  --url http://localhost:8001/services/ \
  --data name=servicepynext \
  --data url=http://host.docker.internal:10084
  
curl --request POST \
  --url http://localhost:8001/services/servicepynext/routes \
  --data paths=/pynext
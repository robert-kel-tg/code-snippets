# Title
Different code snippets

## Dev env:
```bash
docker-compose up
```
## Login into container
```bash
docker-compose exec api sh
```

## Go docs
 godoc -http=":6060"

## Handlers
https://elithrar.github.io/article/testing-http-handlers-go/

## Zipkin
https://github.com/openzipkin/docker-zipkin
### docker port zipkin

```json
curl -vs 127.0.0.1:9411/api/v1/spans -H'Content-type: application/json' -H 'Expect:' -d '[
  {
    "traceId": "5e1b76cb257aa6fd",
    "name": "example - root span",
    "id": "168ba9a2869c3ae1",
    "timestamp": 1473066067938000,
    "duration": 484655,
    "annotations": [
      {
        "timestamp": 1473066067938000,
        "value": "sr",
        "endpoint": {
          "serviceName": "example",
          "ipv4": "0.0.0.0"
        }
      },
      {
        "timestamp": 1473066068422655,
        "value": "ss",
        "endpoint": {
          "serviceName": "example",
          "ipv4": "0.0.0.0"
        }
      }
    ],
    "binaryAnnotations": []
  }
]'
```

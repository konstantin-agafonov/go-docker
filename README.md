# Simple Go + docker setup

to start run the following commands:

```
docker build -t go-docker .
```
```
docker run -p 8080:8080 --rm -v $(pwd):/app -v /app/tmp --name go-docker-air go-docker
```

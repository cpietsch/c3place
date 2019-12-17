# c3place backend

start the redis server

```sh
docker compose up
```

build and start the api server

```sh
make run
```

test the api server

```sh
make test
```

send a pixel with curl

```sh
curl -X POST http://localhost:4000/pixel -H 'Content-Type: application/json' -d '{"r": 0, "g": 255, "b": 0, "x": 100, "y": 100}'
```

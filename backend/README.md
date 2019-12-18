# c3place backend

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

## Environment Variables

| env var       | description               | default     |
|---------------|---------------------------|-------------|
| `PORT`        | http server port          | `4000`      |
| `RATELIMITER` | enable the rate limiter   | `false`     |
| `REDIS_HOST`  | redis host                | `localhost` |
| `REDIS_PORT`  | redis port                | `6379`      |

## Rate Limiter

to enable the rate limiter we need to run a redis server. for that you find a docker-compose.yml file at the root of the repo or if you have a local running redis you can use this one.

run the following command to enable the rate limiter

```sh
RATELIMITER=true ./backend
```

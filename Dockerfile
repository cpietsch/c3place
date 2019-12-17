# build stage
FROM golang:alpine AS build-env
RUN apk --no-cache add build-base git bzr mercurial gcc
ADD . /src
WORKDIR /src/backend
RUN go get -v
RUN go build -o backend

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/backend/backend /app/
ENTRYPOINT ./backend

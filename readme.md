# URL Shortener
### Features

- High perfomance data processing using memcached, MongoDB and GoFiber
- HTTP/2 support
- Load balancing for lower latency and better fault tolerance
- JWT tokenizaton
- Metrics vizualization using Mongo extension for Grafana

### Build
```
docker compose up
```

### API

> **Endpoint:** /,<br> **Method:** POST <br> **Usage example:**
> ```json
> {
>   "name": "sodamntired",
>   "password": "password"
> }
> ```
> **Response example:**
> ```json
> {
>   "jwt": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDc1MTQ2MjQsIm5hbWUiOiJzb2RhbW50aXJlZDEifQ.zEGyXWB5QY0fUCLQ_RGMS2eJwgEs9Z3gvMG_kA6EzcE"
> }
> ```
Returns JWT token or occured error to the user, token expires in 24h.

> **Endpoint:** /,<br> **Method:** PATCH <br> **Usage example:**
> ```json
> {
>   "name": "sodamntired",
>   "password": "password"
> }
> ```
> **Response example:**
> ```json
> {
>   "jwt": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDc1MTQ2MjQsIm5hbWUiOiJzb2RhbW50aXJlZDEifQ.zEGyXWB5QY0fUCLQ_RGMS2eJwgEs9Z3gvMG_kA6EzcE"
> }
> ```
Returns JWT token or occured error to the user, token expires in 24h.

> **Endpoint:** /,<br> **Method:** DELETE <br> **Usage example:**
> ```json
> {
>   "jwt": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDc1MTQ2MjQsIm5hbWUiOiJzb2RhbW50aXJlZDEifQ.zEGyXWB5QY0fUCLQ_RGMS2eJwgEs9Z3gvMG_kA6EzcE",
>   "alias": "WqRi6Y2w"
> }
> ```
> **Response example:**
> ```json
> {
>   "err": null
> }
> ```
Returns an error if occured (e.g. sender is not an author).


> **Endpoint:** /,<br> **Method:** GET <br> **Usage example:**
> ```json
> {
>   "jwt": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDc1MTQ2MjQsIm5hbWUiOiJzb2RhbW50aXJlZDEifQ.zEGyXWB5QY0fUCLQ_RGMS2eJwgEs9Z3gvMG_kA6EzcE",
>   "url": "https://google.com"
> }
> ```
> **Response example:**
> ```json
> {
>   "alias": "WqRi6Y2w"
> }
> ```
Returns an alias or occured error to the user.

> **Endpoint:** /:alias,<br> **Method:** GET

If the alias exists in the database, fetches the url and redirects.




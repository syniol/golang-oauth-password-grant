# OAuth 2.1 Password Grant Type in Golang
![workflow](https://github.com/syniol/golang-oauth2/actions/workflows/makefile.yml/badge.svg)

Implementation of standard OAuth 2.1 for Password Grant type in Golang and its native HTTP server.


## Healthcheck API
```text
GET  oauth2/healthz HTTP/1.1
Host: 127.0.0.1
Content-Type: text/plain
```

__Request:__
```bash
curl -k --location --request GET 'https://127.0.0.1/healthz'
```

__Response:__
Status code `200` (OK) and a simple body response `ok` indicates API is working and operational.
```text
ok
```


## Clients API
This endpoint is responsible for creating a new client/user to be inserted in database.

```text
POST  oauth2/clients HTTP/1.1
Host: 127.0.0.1
Content-Type: application/json
```

__Request:__
```bash
curl -k --location --request POST 'https://127.0.0.1/oauth2/clients' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "johndoe",
    "password": "johnspassword1"
}'
```

__Response:__
```json
{
  "client_id": "a9a6b145-fafe-415c-a92e-c79cbd57567d"
}
```


## Token API
After client registration you can create a token sending a `POST` request to this endpoint.

```text
POST  oauth2/token HTTP/1.1
Host: 127.0.0.1
Content-Type: application/x-www-form-urlencoded
```

__Request:__
```bash
curl -k --location --request POST 'https://127.0.0.1/oauth2/token' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'grant_type=password' \
--data-urlencode 'username=johndoe' \
--data-urlencode 'password=johnspassword1'
```

__Response:__
```json
{
  "access_token": "MmVjZGFiNmY4Y2E2OTQ1ZWNmMGQyNmZlODZhYWM5YzFhNDliYzZiNzNkNmY2MjBmYThiMzM3NTEyODE1ZTc1YjNiZTcxODI3YjFjZDkzZDYyODRkODljZjdjMDU3NWY4M2Y2NjdiODg4ZTliZDIwMzlmMTRlYjkxZGEyYmFkMDM=",
  "token_type": "Bearer",
  "expires_in": 3600
}
```


## Up & Running

```bash
  make deploy
```


### Debug
In order to run debugger you could create a config on your IDE and enable `DEBUG` env variable in your 
local environment. You will need database & cache storage from docker; you could enable them with:

```bash
  make debug
```

![img](https://github.com/syniol/golang-oauth-password-grant/assets/68777073/5c24392a-29df-41c2-8f11-fd32a1053222)

### Todos
 * [ ] Add more documents about this repository and RFC Standard for OAuth 2.1 especially for `password_grant`
 * [ ] Convert Http Error response to JSON response `errors: []`
 * [ ] Investigate possibility of volume share for Redis & Go (app) to share TLS certs
 * [ ] Separate the Docker network for proxy and app to exclude Database (Postgres) & Cache (Redis)
 * [ ] Increase code coverage


#### Credits
Author: [Hadi Tajallaei](mailto:hadi@syniol.com)

Copyright &copy; 2023-2025 Syniol Limited. All rights reserved.

_Please see a [LICENSE file](https://github.com/syniol/golang-oauth-password-grant/blob/main/LICENSE)_

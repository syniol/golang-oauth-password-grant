# OAuth 2.1 Password Grant Type in Golang
![workflow](https://github.com/syniol/golang-oauth-password-grant/actions/workflows/makefile.yml/badge.svg)

Implementation of standard OAuth 2.1 for Password Grant type in Golang 
and its native HTTP server.


## Clients API
```text
POST  oauth2/clients HTTP/1.1
Host: 127.0.0.1:8080
Content-Type: application/json
```

__Request:__
```bash
curl -k --location --request POST 'https://127.0.0.1:8080/oauth2/clients' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "johndoe1",
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
```text
POST  oauth2/token HTTP/1.1
Host: 127.0.0.1:8080
Content-Type: application/x-www-form-urlencoded
```

__Request:__
```bash
curl -k --location --request POST 'https://127.0.0.1:8080/oauth2/token' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'grant_type=password' \
--data-urlencode 'username=johndoe1' \
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
 * [x] SSL For Postgres
 * [x] SSL & Password for Redis
 * [x] Cert for Creation of Token (Could be from Infra or Inside the code)
 * [x] TLS Server Listener
 * [ ] Use Docker Secret to share passwords


#### Credits
Author: [Hadi Tajallaei](mailto:hadi@syniol.com)

Copyright &copy; 2023 Syniol Limited. All rights reserved.

_Please see a [LICENSE file](https://github.com/syniol/golang-oauth-password-grant/blob/main/LICENSE)_

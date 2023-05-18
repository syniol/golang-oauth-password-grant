# OAuth 2 Password Grant Type in Golang
Implementation of standard Oauth V2 for Password Grant type in Golang 
and its native HTTP server.


## Clients API
```text
POST  oauth/clients HTTP/1.1
Host: 127.0.0.1:8080
Content-Type: application/json
```

```bash
curl --location --request POST '127.0.0.1:8080/oauth2/clients' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "johndoe1",
    "password": "johnspassword1"
}'
```

## Token API
```text
POST  oauth/token HTTP/1.1
Host: 127.0.0.1:8080
Content-Type: application/x-www-form-urlencoded
```

```bash
curl --location --request POST '127.0.0.1:8080/oauth2/token' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'grant_type=password' \
--data-urlencode 'username=johndoe1' \
--data-urlencode 'password=johnspassword1'
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
 * [ ] SSL For Postgres
 * [ ] SSL for Redis
 * [ ] Cert for Creation of Token (Could be from Infra or Inside the code)


#### Credits
Author: [Hadi Tajallaei](mailto:hadi@syniol.com)

Copyright &copy; 2023 Syniol Limited. All rights reserved.

_Please see a [LICENSE file](https://github.com/syniol/golang-oauth-password-grant/blob/main/LICENSE)_

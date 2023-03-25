# OAuth Password Grant Type in Golang
Implementation of standard Oauth V2 for Password Grant type.


## Clients API
```text
POST  /clients HTTP/1.1
Host: 127.0.0.1:8080
Content-Type: application/json
```
TODO

## Token API

```text
POST  oauth/token HTTP/1.1
Host: 127.0.0.1:8080
Content-Type: application/x-www-form-urlencoded
```

```bash
curl --location --request POST '127.0.0.1:8080/oauth/token' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'grant_type=password' \
--data-urlencode 'username=johndoe1' \
--data-urlencode 'password=MyPassword!'
```


#### Credits
Author: [Hadi Tajallaei](mailto:hadi@syniol.com)

Copyright &copy; 2023 Syniol Limited. All rights reserved.

# OAuth Password Grant Type in Golang
Implementation of standard Oauth V2 for Password Grant type.


## Token API

```text
POST 127.0.0.1:8080/oauth/token
Host: 127.0.0.1
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

Copyright &copy; 2023. All rights reserved.

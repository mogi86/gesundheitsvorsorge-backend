# How to execute

- Run Server on Docker

```bash
$ make run
```

- Request for Create User

```bash
$ curl -v POST "http://localhost:8000/user/create" \
    -H "Content-Type: application/json" \
    -d '{
        "password": "hogehoge",
        "first_name": "Ken",
        "last_name": "Suzuki",
        "mail": "testhoge@testhoge.com",
        "sex": "1",
        "birthday": "2000-01-01",
        "weight": "60.0",
        "height": "170.0"
    }'
```

- Request for Get User

```bash
$ curl -v GET "http://localhost:8000/user/get?id=3"
```

- Request for Login

```bash
$ curl -v POST "http://localhost:8000/login" \
    -H "Content-Type: application/json" \
    -d '{
        "password": "hogehoge",
        "mail": "testhoge@testhoge.com"
    }'
```

- Request for Home Index

```bash
$ curl -v GET "http://localhost:8000/home/index" \
    -H "Authorization: Bearer your-token"
```
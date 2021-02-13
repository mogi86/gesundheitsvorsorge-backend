# How to execute

- Run Server on Docker

```bash
$ make run
```

- Request

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
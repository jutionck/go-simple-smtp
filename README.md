# Go Simple SMTP

## Config

Create new file `.env` from `.env.example`

```env
EMAIL_SERVER=smtp.gmail.com
EMAIL_PORT=587
EMAIL_FROM=
EMAIL_PASSWORD=
API_HOST=localhost
API_PORT=8888
```

# Run

```bash
go run .
```

# Client

Use Postman

- `POST` -> `/customers
- Body:

```json
{
  "id": "C0001",
  "name": "Jution Candra Kirana",
  "email": "jution.kirana@enigmacamp.com",
  "password": "password"
}
```

- Response:

```json
{
  "message": "Pendafataran sukses"
}
```

And check your mail inbox.

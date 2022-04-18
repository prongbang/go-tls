# go-tls

### Create certificate with mkcert

```shell
mkcert tls.go.com '*.tls.go.com' localhost 127.0.0.1 ::1
```

Logger

```shell

Created a new certificate valid for the following names 📜
 - "tls.go.com"
 - "*.tls.go.com"
 - "localhost"
 - "127.0.0.1"
 - "::1"

Reminder: X.509 wildcards only go one level deep, so this won't match a.b.tls.go.com ℹ️

The certificate is at "./tls.go.com+4.pem" and the key at "./tls.go.com+4-key.pem" ✅

It will expire on 18 July 2024 🗓

```

### Using

```go
cer, err := tls.LoadX509KeyPair("mkcert/tls.go.com+4.pem", "mkcert/tls.go.com+4-key.pem")
```

### Configuration your `/etc/hosts`

- Open vim editor

```shell
sudo vim /etc/hosts
```

- Press key "A" to Edit

- Add this line

```shell
127.0.0.1       tls.go.com
```

- Press key "Esc" to exist edit
- Press key ":wq" to write and quit 


### Run

```shell
go run main.go
```

### Test

```shell
curl https://tls.go.com
```

Result

```json
{
  "protocol": "https"
}
```
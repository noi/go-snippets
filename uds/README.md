# Unix Domain Socket Server and Client

Run Server

```sh
go run . server --unix-socket /tmp/my.sock
```

Run Client

```sh
go run . client --unix-socket /tmp/my.sock http://localhost/hello
```

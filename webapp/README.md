# WebAPP

I learned web application testing.

## Install

```
go get github.com/go-chi/chi/v5
go get github.com/go-chi/chi/v5/middleware
go get github.com/alexedwards/scs/v2
go get github.com/jackc/pgx/v4
go get -u github.com/ory/dockertest/v3
go get github.com/golang-jwt/jwt/v4
```

## User & Password

```
admin@example.com
secret
```

## DockerTest

[GitHub](https://github.com/ory/dockertest)

## go:build コメント
`go:build` コメントを書くとテスト時に以下のように指定してテストのインテグレーションが行えるようになる。

```users_postgres.go
# //とgoの間に隙間は書けない
//go:build integration
```

```
$ go test -v -tags=integration ./...
```

## API Curl

```
$ curl http://localhost:8090/test
{"message":"hello, world"}
```

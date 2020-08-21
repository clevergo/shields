# Shields.io Go SDK for creating an [endpoint](https://shields.io/endpoint).
[![Build Status](https://img.shields.io/travis/clevergo/shields?style=for-the-badge)](https://travis-ci.org/clevergo/shields)
[![Coverage Status](https://img.shields.io/coveralls/github/clevergo/shields?style=for-the-badge)](https://coveralls.io/github/clevergo/shields)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white&style=for-the-badge)](https://pkg.go.dev/clevergo.tech/shields?tab=doc)
[![Go Report Card](https://goreportcard.com/badge/github.com/clevergo/shields?style=for-the-badge)](https://goreportcard.com/report/github.com/clevergo/shields)
[![Release](https://img.shields.io/github/release/clevergo/shields.svg?style=for-the-badge)](https://github.com/clevergo/shields/releases)
[![Downloads](https://img.shields.io/endpoint?url=https://pkg.clevergo.tech/api/badges/downloads/month/clevergo.tech/shields&style=for-the-badge)](https://pkg.clevergo.tech/)

```shell
$ go get -u clevergo.tech/shields
```

```go
import (
	"encoding/json"
	"net/http"

	"clevergo.tech/shields"
)

func handler(w http.ResponseWriter, req *http.Request) {
    badge := shields.New("label", "message")
    badge.LabelColor = shields.ColorBlue
    err := badge.ParseRequest(req)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    data, err := json.Marshal(badge)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(data)
}
```

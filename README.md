# Shields.io Go SDK for creating an [endpoint](https://shields.io/endpoint).
[![Build Status](https://img.shields.io/travis/clevergo/shields?style=flat-square)](https://travis-ci.org/clevergo/shields)
[![Coverage Status](https://img.shields.io/coveralls/github/clevergo/shields?style=flat-square)](https://coveralls.io/github/clevergo/shields)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/clevergo.tech/shields?tab=doc)
[![Go Report Card](https://goreportcard.com/badge/github.com/clevergo/shields?style=flat-square)](https://goreportcard.com/report/github.com/clevergo/shields)
[![Release](https://img.shields.io/github/release/clevergo/shields.svg?style=flat-square)](https://github.com/clevergo/shields/releases)
[![Downloads](https://img.shields.io/endpoint?url=https://pkg.clevergo.tech/api/badges/downloads/total/clevergo.tech/shields&style=flat-square)](https://pkg.clevergo.tech/)
[![Chat](https://img.shields.io/badge/chat-telegram-blue?style=flat-square)](https://t.me/clevergotech)
[![Community](https://img.shields.io/badge/community-forum-blue?style=flat-square&color=orange)](https://forum.clevergo.tech)

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

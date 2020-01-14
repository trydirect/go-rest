package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "io/ioutil"
    "net/http"
    "net/http/httptest"
    "testing"
)

var ts *httptest.Server

func init() {
    gin.SetMode("release")
    engine := GetMainEngine()
    ts = httptest.NewServer(engine)
}

func TestSwaggerIndex(t *testing.T) {
    res, err := http.Get(fmt.Sprintf("%s/swagger/index.html", ts.URL))

    if err != nil {
        t.Fatal(err)
    }

    if res.StatusCode != http.StatusOK {
        t.Errorf("status not OK")
    }

    defer res.Body.Close()
    body, err := ioutil.ReadAll(res.Body)

    if err != nil {
        t.Fatal(err)
    }

    if len(body) == 0 {
        t.Fail()
    }
}

func TestShowAccount(t *testing.T) {
    res, err := http.Get(fmt.Sprintf("%s/api/v1/accounts/1", ts.URL))

    if err != nil {
        t.Fatal(err)
    }

    if res.StatusCode != http.StatusOK {
        t.Errorf("status not OK")
    }

    defer res.Body.Close()
    body, err := ioutil.ReadAll(res.Body)

    if err != nil {
        t.Fatal(err)
    }

    expect := `{"id":1,"name":"account_1","uuid":"00000000-0000-0000-0000-000000000000"}
`
    if string(body) != expect {
        t.Fail()
    }
}
package main

import (
	"api-loja/src/utils"
	"encoding/json"
	"fmt"
	"testing"
)

func TestAuthRouteWithEmptyToken(t *testing.T) {
    _, code, err := utils.Request(utils.RequestConfigInput{
        Method: "GET",
        Url: "http://localhost:8080/ping",
    })
    if err != nil {
        t.Fail()
    }
    if code != 401 {
        t.Fail()
    }
}

func TestAuthWithInvalidToken(t *testing.T) {
    _, code, err := utils.Request(utils.RequestConfigInput{
        Method: "GET",
        Url: "http://127.0.0.1:8080/ping",
        Headers: map[string]string{
            "token":"token invalido",
        },
    })
    if err != nil {
        t.Fail()
    }
    if code != 401 {
        t.Fail()
    }
}
type createTokenRes struct {
    Token string `json:"token"`
}
func TestAuthWithValidToken(t *testing.T) {
    res, code, err := utils.Request(utils.RequestConfigInput{
        Method: "POST",
        Url: "http://loja-auth/create-token",
        Data: "{\"service\":\"api\", \"secret\":\"api-service\"}",
        Headers: map[string]string{
            "Content-type":"application/json",
        },
    })
    if err != nil {
        t.Logf("Erro não é igual a nil 1")
    }
    if code != 200 {
        t.Logf("status de ping é maior que 200 1")
    }
    var resToken createTokenRes
    if err := json.Unmarshal([]byte(res), &resToken); err != nil {
        t.Fail()
    }
    res, code, err = utils.Request(utils.RequestConfigInput{
        Method: "GET",
        Url: "http://localhost:8080/ping",
        Headers: map[string]string{
            "token": resToken.Token,
        },
    })
    if err != nil {
        t.Logf("Erro não é igual a nil 2")
    }
    if code > 299 {
        fmt.Println(res, code, err)
        t.Logf("status de ping é mairo que 200 1")
    }
}

package utils

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
    //testar ping na api do token
    requestConfig := RequestConfigInput{
        Method: "GET",
        Url: "http://loja-auth/ping",
    }
    res, code, err := Request(requestConfig)
    if err != nil {
        t.Fail()
    }
    fmt.Println(res, code)
}

package utils

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)
var client *http.Client
var once sync.Once

// GetClient returns the singleton instance of http.Client
func GetClient() *http.Client {
    once.Do(func() {
        client = &http.Client{
            // Configure your client here if needed, e.g., timeouts
            Timeout: time.Second * 10,
        }
    })
    return client
}
func RequestBody[T any](bodyDest * T, c *gin.Context) bool  {
	if err := c.BindJSON(bodyDest); err != nil {
		log.Println("Erro :: pegando body do request")
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H {
			"mensagem": "corpo de requisição é necessário",
		})
		return false
	}
	return true
}
type RequestConfigInput struct {
    Method string
    Url string
    Data string
    DataType string
    Headers map[string]string
}
func Request(config RequestConfigInput) (string, int, error) {
    code := 0
    body_ := ""
	// Verificar se os parâmetros obrigatórios estão presentes
	if config.Method == "" || config.Url == "" {
		fmt.Println("Method, Url, and Data are required.")
		return body_, code, errors.New("Method, Url, and Data are required.")
	}

	// Preparar o corpo da requisição de acordo com o DataType
	var body *bytes.Buffer
	if strings.ToLower(config.DataType) == "json" {
		body = bytes.NewBuffer([]byte(config.Data))
	} else if strings.ToLower(config.DataType) == "query" {
		data := url.Values{}
		pairs := strings.Split(config.Data, "&")
		for _, pair := range pairs {
			kv := strings.Split(pair, "=")
			if len(kv) == 2 {
				data.Set(kv[0], kv[1])
			}
		}
		body = bytes.NewBufferString(data.Encode())
	} else {
		// Se o DataType não for especificado, assume-se JSON como padrão
		body = bytes.NewBuffer([]byte(config.Data))
	}

	// Criar a requisição HTTP
	req, err := http.NewRequest(config.Method, config.Url, body)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return  body_, code, err
	}

	// Adicionar cabeçalhos (headers) se existirem
	if config.Headers != nil {
		for key, value := range config.Headers {
			req.Header.Set(key, value)
		}
	}

	// Se o DataType for Json, adicionar o cabeçalho Content-Type como application/json
	if strings.ToLower(config.DataType) == "json" {
		req.Header.Set("Content-Type", "application/json")
	}

	// Fazer a requisição
	client := GetClient()
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return  body_, code, err
	}
	defer resp.Body.Close()

	// Ler a resposta
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return  body_, code, err
	}

    return string(bodyBytes), resp.StatusCode, nil
}

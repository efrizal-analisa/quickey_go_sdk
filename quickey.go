package quickey

// "github.com/efrizal-analisa/quickey_go_sdk/app"
// "github.com/efrizal-analisa/quickey_go_sdk/auth"
import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	APIVersion = "v0.1.0"
	APIURL     = "https://api.getquickey.com"
)

type Response struct {
	ApiKey  string
	BaseUrl string
}

type App struct {
	Data interface{} `json:"app"`
}

type Auth struct {
	Token interface{} `json:"access_token"`
}

func New(api_key string) *Response {
	return &Response{
		ApiKey:  api_key,
		BaseUrl: APIURL,
	}
}

func (q *Response) GetMetadata() *App {
	values := map[string]string{"api_key": q.ApiKey}
	json_data, err := json.Marshal(values)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(q.BaseUrl+"/auth/apiKey", "application/json",
		bytes.NewBuffer(json_data))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)

	// var res map[string]interface{}
	// json.NewDecoder(resp.Body).Decode()
	return &App{}
}

// func (q *Response) GetAccessToken() *Auth {
// }

package quickey

// "github.com/efrizal-analisa/quickey_go_sdk/app"
// "github.com/efrizal-analisa/quickey_go_sdk/auth"
import (
	"bytes"
	"encoding/json"

	// "fmt"
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

type Data struct {
	Email          string `json:"email"`
	AppName        string `json:"appName"`
	SocialApps     string `json:"socialApps"`
	RedirectUri    string `json:"redirectUri"`
	RedirectUrlApp string `json:"redirectUrlApp"`
	ApiKey         string `json:"apiKey"`
}

type App struct {
	Data Data `json:"app"`
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

func (q *Response) GetMetadata(w http.ResponseWriter, r *http.Request) string {
	// var data Data
	// renderResponse(w,
	// 	&App{
	// 		Data: Data{
	// 			Email:          data.Email,
	// 			AppName:        data.AppName,
	// 			SocialApps:     data.SocialApps,
	// 			RedirectUri:    data.RedirectUri,
	// 			RedirectUrlApp: data.RedirectUrlApp,
	// 			ApiKey:         data.ApiKey,
	// 		},
	// 	},
	// 	http.StatusCreated)

	values := map[string]string{"api_key": q.ApiKey}
	json_data, err := json.Marshal(values)

	if err != nil {
		log.Fatal(err)
	}

	responseJSON, err := http.Post(q.BaseUrl+"/auth/apiKey", "application/json",
		bytes.NewBuffer(json_data))

	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	responseBytes, err := json.Marshal(responseJSON)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)

	if _, err = w.Write(responseBytes); err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	return string(responseBytes)
}

// func (q *Response) GetAccessToken() *Auth {
// }

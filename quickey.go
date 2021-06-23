package quickey

import (
	"bytes"
	"encoding/json"
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
	App     []App
	Auth    []Auth
}

type App struct {
	Email          string `json:"email"`
	AppName        string `json:"appName"`
	SocialApps     string `json:"socialApps"`
	RedirectUri    string `json:"redirectUri"`
	RedirectUrlApp string `json:"redirectUrlApp"`
	ApiKey         string `json:"apiKey"`
}

type Auth struct {
	Token string `json:"access_token"`
}

func New(api_key string) *Response {
	return &Response{
		ApiKey:  api_key,
		BaseUrl: APIURL,
	}
}

func (q *Response) GetMetadata() *App {

	values := map[string]string{"apiKey": q.ApiKey}
	json_data, err := json.Marshal(values)
	if err != nil {
		log.Fatal(err)
	}

	responseJSON, err := http.Post(q.BaseUrl+"/auth/apiKey", "application/json",
		bytes.NewBuffer(json_data))

	if err != nil {
		log.Fatal(err)
	}

	var responseMap map[string]interface{}
	json.NewDecoder(responseJSON.Body).Decode(&responseMap)

	responseBytes, err := json.Marshal(responseMap["app"])

	responseString := string(responseBytes)

	app := App{}
	json.Unmarshal([]byte(responseString), &app)
	q.App = append(q.App, app)

	return &app
}

func (q *Response) GetAccessToken(email string) *Auth {

	values := map[string]string{"email": email}
	json_data, err := json.Marshal(values)
	if err != nil {
		log.Fatal(err)
	}

	responseJSON, err := http.Post(q.BaseUrl+"/loginRegister", "application/json",
		bytes.NewBuffer(json_data))

	if err != nil {
		log.Fatal(err)
	}

	var responseMap map[string]interface{}
	json.NewDecoder(responseJSON.Body).Decode(&responseMap)

	responseBytes, err := json.Marshal(responseMap)

	responseString := string(responseBytes)

	auth := Auth{}
	json.Unmarshal([]byte(responseString), &auth)
	q.Auth = append(q.Auth, auth)

	return &auth
}

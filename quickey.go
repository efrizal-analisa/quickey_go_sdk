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

func GetMetadata() App {
	// values := map[string]string{"api_key": q.ApiKey}
	values := map[string]string{"apiKey": "inT9Ic-BhfqbRA-wgtz8Dn_WHUuAAmSI3VN0kByQpyU"}

	json_data, err := json.Marshal(values)

	if err != nil {
		log.Fatal(err)
	}

	// responseJSON, err := http.Post(q.BaseUrl+"/auth/apiKey", "application/json",
	// 	bytes.NewBuffer(json_data))

	responseJSON, err := http.Post("https://api.getquickey.com/auth/apiKey", "application/json",
		bytes.NewBuffer(json_data))

	if err != nil {
		log.Fatal(err)
	}
	var responseMap map[string]interface{}
	// var res["app"] map[string]interface{}

	json.NewDecoder(responseJSON.Body).Decode(&responseMap)

	responseBytes, err := json.Marshal(responseMap["app"])
	responseString := string(responseBytes)
	// fmt.Println(jsonString)
	app := App{}
	json.Unmarshal([]byte(responseString), &app)
	return app
}

// func (q *Response) GetAccessToken() *Auth {
// }

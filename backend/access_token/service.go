package access_token

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"hello-world/env"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

var accessToken = ""

func InitAccessToken() error {

	baseUrl := fmt.Sprintf("https://zoom.us/oauth/token?grant_type=account_credentials&account_id=%s", env.ApplicationId)

	account := env.ClientId + ":" + env.ClientSecret
	authorization := base64.StdEncoding.EncodeToString([]byte(account))

	client := &http.Client{}

	req, err := http.NewRequest("POST", baseUrl, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "Basic "+authorization)
	response, err := client.Do(req)

	if err != nil {
		return err
	}

	post := &AccessTokenResponse{}
	err = json.NewDecoder(response.Body).Decode(post)

	if err != nil {

		return err
	}

	fmt.Println("Get new access token", post.AccessToken)

	accessToken = post.AccessToken

	return nil
}

func SetAccessToken(request *http.Request) {
	request.Header.Add("Content-Type", "application/json")

	CheckAccessToken()

	request.Header.Add("Authorization", "Bearer "+accessToken)

}

func CheckAccessToken() {
	token, _, err := new(jwt.Parser).ParseUnverified(accessToken, jwt.MapClaims{})
	if err != nil {
		fmt.Println(err)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		fmt.Println(claims["foo"], claims["exp"])
	} else {
		fmt.Println(err)
	}
}

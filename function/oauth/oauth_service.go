package oauth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/go-retryablehttp"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strconv"

	"github.com/mattermost/mattermost-plugin-apps/apps"
)

type OauthService interface {
	RefreshToken() Token
}

type OauthServiceImpl struct {
	Creq apps.CallRequest
}

func (s OauthServiceImpl) RefreshToken() Token {

	clientId := s.Creq.Context.OAuth2.OAuth2App.ClientID
	clientSecret := s.Creq.Context.OAuth2.OAuth2App.ClientSecret
	remoteUrl := s.Creq.Context.OAuth2.OAuth2App.RemoteRootURL

	reqUrl := fmt.Sprintf("%s/index.php/apps/oauth2/api/v1/token", remoteUrl)
	refreshToken := s.Creq.Context.OAuth2.User.(map[string]interface{})["refresh_token"].(string)

	payload := RefreshTokenBody{
		RefreshToken: refreshToken,
		GrantType:    "refresh_token",
	}

	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", reqUrl, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.SetBasicAuth(clientId, clientSecret)
	maxRetries, _ := strconv.Atoi(os.Getenv("MAX_REQUEST_RETRIES"))
	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = maxRetries

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Errorf("Error during refreshing of the token. Error: %s", err)
	}

	jsonResp := Token{}
	json.NewDecoder(resp.Body).Decode(&jsonResp)

	return jsonResp

}

func getToken(creq apps.CallRequest) Token {
	code, _ := creq.Values["code"].(string)

	clientId := creq.Context.OAuth2.OAuth2App.ClientID
	clientSecret := creq.Context.OAuth2.OAuth2App.ClientSecret
	remoteUrl := creq.Context.OAuth2.OAuth2App.RemoteRootURL

	reqUrl := fmt.Sprintf("%s/index.php/apps/oauth2/api/v1/token", remoteUrl)

	payload := RequestTokenBody{
		Code:      code,
		GrantType: "authorization_code",
	}

	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", reqUrl, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.SetBasicAuth(clientId, clientSecret)

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Errorf("Error during getting of the token. Error: %s", err)
	}
	jsonResp := Token{}
	json.NewDecoder(resp.Body).Decode(&jsonResp)
	return jsonResp
}

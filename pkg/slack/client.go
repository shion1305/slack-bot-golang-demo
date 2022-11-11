package slack

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

const (
	ApiAccessTokenEndpoint   string = "https://slack.com/api/oauth.v2.access"
	ApiChannelListEndpoint   string = "https://slack.com/api/conversations.list"
	ApiChannelCreateEndpoint string = "https://slack.com/api/conversations.create"
)

type SlackAPI struct {
	clientId     string
	clientSecret string
	redirectURI  string
}

func NewSlackAPI(
	clientId string, clientSecret string, redirectUri string,
) SlackAPI {
	return SlackAPI{
		clientId:     clientId,
		clientSecret: clientSecret,
		redirectURI:  redirectUri}
}

// GetAccessToken refer https://api.slack.com/methods/oauth.v2.access
func (api SlackAPI) GetAccessToken(code string) (*AccessResponse, error) {
	client := resty.New()
	var r AccessResponse
	resp, err := client.R().
		SetQueryParams(map[string]string{
			"code":         code,
			"redirect_uri": api.redirectURI,
		}).
		SetBasicAuth(api.clientId, api.clientSecret).
		SetResult(&r).
		Post(ApiAccessTokenEndpoint)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		err := fmt.Errorf(
			"error status %d: %s",
			resp.StatusCode(),
			resp.Body(),
		)
		return nil, err
	}
	return &r, nil
}

// GetChannelList refer https://api.slack.com/methods/conversations.list
func (api SlackAPI) GetChannelList(accessToken string) {

}

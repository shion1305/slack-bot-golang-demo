package slack

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

const (
	ApiAccessTokenEndpoint      string = "https://slack.com/api/oauth.v2.access"
	ApiConversationListEndpoint string = "https://slack.com/api/conversations.list"
	ApiChannelCreateEndpoint    string = "https://slack.com/api/conversations.create"
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

// GetConversationList refer https://api.slack.com/methods/conversations.list
func (api SlackAPI) GetConversationList() (*ConversationListResponse, error) {
	client := resty.New()
	var r ConversationListResponse
	resp, err := client.R().
		SetQueryParams(map[string]string{
			"limit":            "1000",
			"exclude_archived": "false",
			"types":            "public_channel",
		}).
		SetBasicAuth(api.clientId, api.clientSecret).
		SetResult(&r).
		Get(ApiConversationListEndpoint)
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

//func handleResp(resp *resty.Response, r *ResponseStatus, err error) (*interface{}, error) {
//	if err != nil {
//		return nil, err
//	}
//
//	if resp.StatusCode() != 200 {
//		err := fmt.Errorf(
//			"error status %d: %s",
//			resp.StatusCode(),
//			resp.Body(),
//		)
//		return nil, err
//	}
//	return r, nil
//}

package slack

import (
	slackLib "github.com/slack-go/slack"
	"net/http"
)

type SlackAPI struct {
	clientId     string
	clientSecret string
	redirectURI  string
	token        string
	httpClient   *http.Client
	slackClient  *slackLib.Client
}

func NewSlackAPI(
	clientId string, clientSecret string, redirectUri string, token string,
) SlackAPI {

	return SlackAPI{
		clientId:     clientId,
		clientSecret: clientSecret,
		redirectURI:  redirectUri,
		token:        token,
		httpClient:   &http.Client{},
		slackClient:  slackLib.New(token),
	}
}

// GetAccessToken refer https://api.slack.com/methods/oauth.v2.access
func (api SlackAPI) GetAccessToken(code string) (*slackLib.OAuthV2Response, error) {
	return slackLib.GetOAuthV2Response(api.httpClient, api.clientId, api.clientSecret, code, api.redirectURI)
}

// GetConversationList refer https://api.slack.com/methods/conversations.list
func (api SlackAPI) GetConversationList() ([]slackLib.Channel, error) {
	channels, nextCursor, err := api.slackClient.GetConversations(
		&slackLib.GetConversationsParameters{
			Types: []string{"public_channel", "private_channel"},
		},
	)
	for nextCursor != "" {
		nextChannels, c, err := api.slackClient.GetConversations(
			&slackLib.GetConversationsParameters{
				Types:  []string{"public_channel", "private_channel"},
				Cursor: nextCursor,
			},
		)
		if err != nil {
			return nil, err
		}
		nextCursor = c
		channels = append(channels, nextChannels...)
	}
	return channels, err
}

func (api SlackAPI) CreateConversation(conversationName string, isPrivate bool) (*slackLib.Channel, error) {
	return api.slackClient.CreateConversation(conversationName, isPrivate)
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

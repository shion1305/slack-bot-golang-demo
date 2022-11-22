package slack

import (
	slackLib "github.com/slack-go/slack"
	"net/http"
)

type SlackAPI struct {
	clientId     string
	clientSecret string
	redirectURI  string
	httpClient   *http.Client
}

func NewSlackAPI(
	clientId string, clientSecret string, redirectUri string,
) SlackAPI {

	return SlackAPI{
		clientId:     clientId,
		clientSecret: clientSecret,
		redirectURI:  redirectUri,
		httpClient:   &http.Client{},
	}
}

// GetAccessToken refer https://api.slack.com/methods/oauth.v2.access
func (api SlackAPI) GetAccessToken(code string) (*slackLib.OAuthV2Response, error) {
	return slackLib.GetOAuthV2Response(api.httpClient, api.clientId, api.clientSecret, code, api.redirectURI)
}

// GetConversationList refer https://api.slack.com/methods/conversations.list
func (api SlackAPI) GetConversationList(token string) ([]slackLib.Channel, error) {
	slackClient := slackLib.New(token)
	channels, nextCursor, err := slackClient.GetConversations(
		&slackLib.GetConversationsParameters{
			Types: []string{"public_channel", "private_channel"},
		},
	)
	for nextCursor != "" {
		nextChannels, c, err := slackClient.GetConversations(
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

func (api SlackAPI) CreateConversation(token string, conversationName string, isPrivate bool) (*slackLib.Channel, error) {
	return slackLib.New(token).CreateConversation(conversationName, isPrivate)
}

func (api SlackAPI) PostMessage(token string, channel string, message string) (conversationID string, timestampID string, err error) {
	return slackLib.New(token).PostMessage(channel, slackLib.MsgOptionText(message, false))
}

func (api SlackAPI) PostMessageWithAttachment(token string, channel string, attachment slackLib.Attachment) (conversationID string, timestampID string, err error) {
	return slackLib.New(token).PostMessage(channel, slackLib.MsgOptionAttachments(attachment))
}

func (api SlackAPI) GetUsers(token string) ([]slackLib.User, error) {
	return slackLib.New(token).GetUsers()
}

func (api SlackAPI) CloseConversation(token string, channel string) (noOp bool, alreadyClosed bool, error error) {
	return slackLib.New(token).CloseConversation(channel)
}

func (api SlackAPI) InviteUsersToConversation(token string, channel string, users []string) (*slackLib.Channel, error) {
	return slackLib.New(token).InviteUsersToConversation(channel, users...)
}

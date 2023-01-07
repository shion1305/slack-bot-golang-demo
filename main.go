package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"goSlackBotEmpath/pkg/slack"
	"net/url"
	"os"
)

const (
	EnvClientID     = "clientId"
	EnvClientSecret = "clientSecret"
	EnvRedirectURI  = "redirectURI"
	EnvHost         = "host"
	EnvPort         = "port"
)

func genAuthURI() string {
	return "https://slack.com/oauth/v2/authorize?client_id=" + os.Getenv("clientId") + "&scope=channels:join,channels:manage,chat:write,channels:read,groups:read,users:read,users.profile:read&redirect_uri=" + url.QueryEscape(os.Getenv(EnvRedirectURI)) + "&state=123"
}

var api slack.SlackAPI

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Errorf("error loading .env file: %v", err))
		return
	}
	api = slack.NewSlackAPI(os.Getenv(EnvClientID), os.Getenv(EnvClientSecret), os.Getenv(EnvRedirectURI))
	client := gin.Default()

	client.GET("/slack/auth", devAuth)
	client.GET("/callback", authCallback)
	fmt.Println(os.Getenv(EnvHost) + "/slack/auth")
	err = client.Run(":" + os.Getenv(EnvPort))
	if err != nil {
		panic(err)
	}
}

func devAuth(c *gin.Context) {
	c.Redirect(302, genAuthURI())
}

func authCallback(c *gin.Context) {
	code := c.Query("code")
	fmt.Println(code)
	accessToken, err := api.GetAccessToken(code)
	if err != nil {
		_, _ = c.Writer.WriteString("Error: " + err.Error() + "\n")
		return
	}
	_, _ = c.Writer.WriteString("AccessToken: " + accessToken.AccessToken + "\n\n")

	testGetUserInfo(accessToken.AccessToken, accessToken.AuthedUser.ID, c)

	profile := testGetUserProfile(accessToken.AccessToken, c)

	//try CreateConversation
	channelID := testCreateConversation(accessToken.AccessToken, c)

	//try GetConversationList
	testGetConversationList(accessToken.AccessToken, c)

	_, _ = c.Writer.WriteString("\n")

	testGetUsers(accessToken.AccessToken, c)

	_, _ = c.Writer.WriteString("\n")

	msg := fmt.Sprintf("Hello %s, your user id is %s", profile.DisplayName, accessToken.AuthedUser.ID)
	testPostMessage(accessToken.AccessToken, channelID, msg, c)

	_, _ = c.Writer.WriteString("\n")

	mentions := []string{accessToken.AuthedUser.ID}

	testSendNotificationTemplate(accessToken.AccessToken, channelID, mentions, c)

	testSendDirectMessage(accessToken.AccessToken, accessToken.AuthedUser.ID, c)
}

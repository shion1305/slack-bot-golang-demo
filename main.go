package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"goSlackBotEmpath/pkg/slack"
	"net/url"
	"os"
	"time"
)

const (
	EnvClientID     = "clientId"
	EnvClientSecret = "clientSecret"
	EnvRedirectURI  = "redirectURI"
	EnvHost         = "host"
	EnvPort         = "port"
)

func genAuthURI() string {
	return "https://slack.com/oauth/v2/authorize?client_id=" + os.Getenv("clientId") + "&scope=channels:join,channels:manage,chat:write,channels:read,groups:read&redirect_uri=" + url.QueryEscape(os.Getenv(EnvRedirectURI)) + "&state=123"
}

var api slack.SlackAPI

func main() {
	godotenv.Load()
	api = slack.NewSlackAPI(os.Getenv(EnvClientID), os.Getenv(EnvClientSecret), os.Getenv(EnvRedirectURI))
	client := gin.Default()

	client.GET("/slack/auth", devAuth)
	client.GET("/callback", authCallback)
	fmt.Println(os.Getenv(EnvHost) + "/slack/auth")
	client.Run(":" + os.Getenv(EnvPort))
}

func devAuth(c *gin.Context) {
	c.Redirect(302, genAuthURI())
}

func authCallback(c *gin.Context) {
	code := c.Query("code")
	fmt.Println(code)
	accessToken, err := api.GetAccessToken(code)
	if err != nil {
		c.Writer.WriteString("Error: " + err.Error() + "\n")
		return
	}
	c.Writer.WriteString("AccessToken: " + accessToken.AccessToken + "\n\n")

	//try CreateConversation
	channelID := testCreateConversation(accessToken.AccessToken, c)

	//try GetConversationList
	testGetConversationList(accessToken.AccessToken, c)

	c.Writer.WriteString("\n")

	testGetUsers(accessToken.AccessToken, c)

	c.Writer.WriteString("\n")

	testPostMessage(accessToken.AccessToken, channelID, c)

	c.Writer.WriteString("\n")

	c.Writer.WriteString("CloseConversation count 10...\n")
	time.Sleep(10 * time.Second)

	testCloseConversation(accessToken.AccessToken, channelID, c)
}

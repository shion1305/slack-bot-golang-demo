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

func genAuthURI() {
	//api := slack.NewSlackAPI(os.Getenv("clientId"), os.Getenv("clientSecret"), os.Getenv("redirectURI"))
	fmt.Println("https://slack.com/oauth/v2/authorize?client_id=" + os.Getenv(EnvClientID) + "&scope=channels:join,channels:manage,chat:write,channels:read,groups:read&redirect_uri=" + os.Getenv(EnvRedirectURI))

}

func main() {
	godotenv.Load()
	api := slack.NewSlackAPI(os.Getenv(EnvClientID), os.Getenv(EnvClientSecret), os.Getenv(EnvRedirectURI))

	client := gin.Default()

	client.GET("/slack/auth", func(c *gin.Context) {
		authUrl := "https://slack.com/oauth/v2/authorize?client_id=" + os.Getenv("clientId") + "&scope=channels:join,channels:manage,chat:write,channels:read,groups:read&redirect_uri=" + url.QueryEscape(os.Getenv(EnvRedirectURI)) + "&state=123"
		fmt.Println(authUrl)
		c.Redirect(302, authUrl)
	})

	client.GET("/callback", func(c *gin.Context) {
		code := c.Query("code")
		fmt.Println(code)
		accessToken, err := api.GetAccessToken(code)
		if err != nil {
			c.Writer.WriteString("Error: " + err.Error() + "\n")
			return
		}
		c.Writer.WriteString("AccessToken: " + accessToken.AccessToken + "\n\n")

		//try CreateConversation
		channel, err := api.CreateConversation(accessToken.AccessToken, "test", false)
		if err != nil {
			c.Writer.WriteString("Error: " + err.Error() + "\n")
			return
		}
		c.Writer.WriteString("Created Channel!!!\n")
		c.Writer.WriteString("ChannelID: " + channel.ID + "\n")
		c.Writer.WriteString("ChannelName: " + channel.Name + "\n")

		//try GetConversationList
		conversationList, err := api.GetConversationList(accessToken.AccessToken)
		if err != nil {
			c.Writer.WriteString("Error in ConversationList: " + err.Error() + "\n")
			return
		}
		c.Writer.WriteString("ConversationList: " + "\n")
		for _, conversation := range conversationList {
			// Output Name, ID, IsChannel, Creator, Members
			c.Writer.WriteString("Name: " + conversation.Name + "\n")
			c.Writer.WriteString("ID: " + conversation.ID + "\n")
			c.Writer.WriteString("IsChannel: " + fmt.Sprintf("%v", conversation.IsChannel) + "\n")
			c.Writer.WriteString("Creator: " + conversation.Creator + "\n")
			c.Writer.WriteString("Members: " + fmt.Sprintf("%v", conversation.Members) + "\n")
			c.Writer.WriteString("\n")
		}

		//try PostMessage
		t, t1, err := api.PostMessage(accessToken.AccessToken, channel.ID, "Hello World!")
		if err != nil {
			c.Writer.WriteString("Error in PostMessage: " + err.Error() + "\n")
			return
		}
		c.Writer.WriteString("PostMessage: " + "\n")
		c.Writer.WriteString("Timestamp: " + t + "\n")
		c.Writer.WriteString("Message: " + t1 + "\n")
	})
	fmt.Println(os.Getenv(EnvHost) + "/slack/auth")
	client.Run(":" + os.Getenv(EnvPort))
}

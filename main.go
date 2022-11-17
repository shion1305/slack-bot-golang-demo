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
	EnvToken        = "token"
)

func genAuthURI() {
	//api := slack.NewSlackAPI(os.Getenv("clientId"), os.Getenv("clientSecret"), os.Getenv("redirectURI"))
	fmt.Println("https://slack.com/oauth/v2/authorize?client_id=" + os.Getenv(EnvClientID) + "&scope=channels:join,channels:manage,chat:write,channels:read,groups:read&redirect_uri=" + os.Getenv(EnvRedirectURI))

}

func main() {
	godotenv.Load()
	api := slack.NewSlackAPI(os.Getenv(EnvClientID), os.Getenv(EnvClientSecret), os.Getenv(EnvRedirectURI), os.Getenv(EnvToken))

	client := gin.Default()

	client.GET("/slack/auth", func(c *gin.Context) {
		authUrl := "https://slack.com/oauth/v2/authorize?client_id=" + os.Getenv("clientId") + "&scope=channels:join,channels:manage,chat:write,channels:read,groups:read&redirect_uri=" + url.QueryEscape(os.Getenv(EnvRedirectURI)) + "&state=123"
		fmt.Println(authUrl)
		c.Redirect(302, authUrl)
	})

	client.GET("/callback", func(c *gin.Context) {
		code := c.Query("code")
		fmt.Println(code)
		_, err := api.GetAccessToken(code)
		if err != nil {
			c.Writer.WriteString("Error1: " + err.Error() + "\n")
			return
		}
		accessToken, err := api.GetAccessToken(code)
		if err != nil {
			c.Writer.WriteString("Error2: " + err.Error() + "\n")
			c.Writer.WriteString(accessToken.AccessToken + "\n")
			return
		}
		c.Writer.WriteString("AccessToken: " + accessToken.AccessToken + "\n")
	})
	fmt.Println(os.Getenv(EnvHost) + "/slack/auth")
	client.Run(":" + os.Getenv(EnvPort))
}

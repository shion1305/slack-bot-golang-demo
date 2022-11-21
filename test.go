package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func testCreateConversation(token string, c *gin.Context) string {
	// get current timestamp as string
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	channel, err := api.CreateConversation(token, "test-"+timestamp, false)
	if err != nil {
		c.Writer.WriteString("Error: " + err.Error() + "\n")
		return ""
	}
	c.Writer.WriteString("---Created Channel---\n")
	c.Writer.WriteString("ChannelID: " + channel.ID + "\n")
	c.Writer.WriteString("ChannelName: " + channel.Name + "\n")
	api.InviteUsersToConversation(token, channel.ID, []string{"U04936U1UEB"})
	return channel.ID
}

func testGetConversationList(token string, c *gin.Context) {
	conversationList, err := api.GetConversationList(token)
	if err != nil {
		c.Writer.WriteString("Error in ConversationList: " + err.Error() + "\n")
		return
	}
	c.Writer.WriteString("---ConversationList---" + "\n")
	for _, conversation := range conversationList {
		// Output Name, ID, IsChannel, Creator, Members
		c.Writer.WriteString("Name: " + conversation.Name + "\n")
		c.Writer.WriteString("ID: " + conversation.ID + "\n")
		c.Writer.WriteString("IsChannel: " + fmt.Sprintf("%v", conversation.IsChannel) + "\n")
		c.Writer.WriteString("Creator: " + conversation.Creator + "\n")
		c.Writer.WriteString("Members: " + fmt.Sprintf("%v", conversation.Members) + "\n")
		c.Writer.WriteString("\n")
	}
}

func testPostMessage(token string, channelID string, c *gin.Context) {
	if channelID == "" {
		return
	}
	t, t1, err := api.PostMessage(token, channelID, "Hello World!")
	if err != nil {
		c.Writer.WriteString("Error in PostMessage: " + err.Error() + "\n")
		return
	}
	c.Writer.WriteString("---PostMessage---" + "\n")
	c.Writer.WriteString("ConversationID: " + t + "\n")
	c.Writer.WriteString("TimestampID: " + t1 + "\n")
}

func testGetUsers(token string, c *gin.Context) []string {
	users, err := api.GetUsers(token)
	if err != nil {
		c.Writer.WriteString("Error in GetUsers: " + err.Error() + "\n")
		return nil
	}
	c.Writer.WriteString("---GetUsers---" + "\n")
	var result []string
	for _, user := range users {
		// Output ID, Name, isBot, Profile.Email, Profile.DisplayName, Profile.DisplayNameNormalized, Profile.RealName, Profile.RealNameNormalized, Profile.Image192, Profile.Image512
		c.Writer.WriteString("ID: " + user.ID + "\n")
		c.Writer.WriteString("Name: " + user.Name + "\n")
		c.Writer.WriteString("isBot: " + fmt.Sprintf("%v", user.IsBot) + "\n")
		c.Writer.WriteString("Profile.Email: " + user.Profile.Email + "\n")
		c.Writer.WriteString("Profile.DisplayName: " + user.Profile.DisplayName + "\n")
		c.Writer.WriteString("Profile.DisplayNameNormalized: " + user.Profile.DisplayNameNormalized + "\n")
		c.Writer.WriteString("Profile.RealName: " + user.Profile.RealName + "\n")
		c.Writer.WriteString("Profile.RealNameNormalized: " + user.Profile.RealNameNormalized + "\n")
		c.Writer.WriteString("Profile.Image192: " + user.Profile.Image192 + "\n")
		c.Writer.WriteString("Profile.Image512: " + user.Profile.Image512 + "\n")
		c.Writer.WriteString("\n")
		result = append(result, user.ID)
	}
	c.Writer.WriteString(fmt.Sprintf("%+q", result))
	c.Writer.WriteString("\n")
	return result
}

func testCloseConversation(token string, channelID string, c *gin.Context) {
	if channelID == "" {
		return
	}
	noOp, alreadyClosed, err := api.CloseConversation(token, channelID)

	//print result to c
	c.Writer.WriteString("---CloseConversation---" + "\n")
	c.Writer.WriteString("noOp: " + fmt.Sprintf("%v", noOp) + "\n")
	c.Writer.WriteString("alreadyClosed: " + fmt.Sprintf("%v", alreadyClosed) + "\n")
	c.Writer.WriteString("err: " + fmt.Sprintf("%v", err) + "\n")

	if err != nil {
		c.Writer.WriteString("Error in CloseConversation: " + err.Error() + "\n")
		return
	}
}

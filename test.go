package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func testCreateConversation(token string, c *gin.Context) string {
	channel, err := api.CreateConversation(token, "test", false)
	if err != nil {
		c.Writer.WriteString("Error: " + err.Error() + "\n")
		return ""
	}
	c.Writer.WriteString("---Created Channel---\n")
	c.Writer.WriteString("ChannelID: " + channel.ID + "\n")
	c.Writer.WriteString("ChannelName: " + channel.Name + "\n")
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

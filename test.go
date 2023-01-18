package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/slack-go/slack"
	uc "goSlackBotEmpath/uc/slack"
	"strconv"
	"time"
)

func testCreateConversation(token string, c *gin.Context) string {
	// get current timestamp as string
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	channel, err := api.CreateConversation(token, "test-"+timestamp, false)
	if err != nil {
		_, _ = c.Writer.WriteString("Error: " + err.Error() + "\n")
		return ""
	}
	_, _ = c.Writer.WriteString("---Created Channel---\n")
	_, _ = c.Writer.WriteString("ChannelID: " + channel.ID + "\n")
	_, _ = c.Writer.WriteString("ChannelName: " + channel.Name + "\n")
	api.InviteUsersToConversation(token, channel.ID, []string{"U04936U1UEB"})
	return channel.ID
}

func testGetConversationList(token string, c *gin.Context) {
	conversationList, err := api.GetConversationList(token)
	if err != nil {
		_, _ = c.Writer.WriteString("Error in ConversationList: " + err.Error() + "\n")
		return
	}
	_, _ = c.Writer.WriteString("---ConversationList---" + "\n")
	for _, conversation := range conversationList {
		// Output Name, ID, IsChannel, Creator, Members
		_, _ = c.Writer.WriteString("Name: " + conversation.Name + "\n")
		_, _ = c.Writer.WriteString("ID: " + conversation.ID + "\n")
		_, _ = c.Writer.WriteString("IsChannel: " + fmt.Sprintf("%v", conversation.IsChannel) + "\n")
		_, _ = c.Writer.WriteString("Creator: " + conversation.Creator + "\n")
		_, _ = c.Writer.WriteString("Members: " + fmt.Sprintf("%v", conversation.Members) + "\n")
		_, _ = c.Writer.WriteString("IsOpen: " + fmt.Sprintf("%v", conversation.IsOpen) + "\n")
		_, _ = c.Writer.WriteString("IsPrivate: " + fmt.Sprintf("%v", conversation.IsPrivate) + "\n")
		_, _ = c.Writer.WriteString("IsArchived: " + fmt.Sprintf("%v", conversation.IsArchived) + "\n")
		_, _ = c.Writer.WriteString("\n")
	}
}

func testPostMessage(token string, channelID string, msg string, c *gin.Context) {
	if channelID == "" {
		return
	}
	t, t1, err := api.PostMessage(token, channelID, msg)
	if err != nil {
		_, _ = c.Writer.WriteString("Error in PostMessage: " + err.Error() + "\n")
		return
	}
	_, _ = c.Writer.WriteString("---PostMessage---" + "\n")
	_, _ = c.Writer.WriteString("ConversationID: " + t + "\n")
	_, _ = c.Writer.WriteString("TimestampID: " + t1 + "\n")
}

func testGetUsers(token string, c *gin.Context) []string {
	users, err := api.GetUsers(token)
	if err != nil {
		_, _ = c.Writer.WriteString("Error in GetUsers: " + err.Error() + "\n")
		return nil
	}
	_, _ = c.Writer.WriteString("---GetUsers---" + "\n")
	var result []string
	for _, user := range users {
		// Output ID, Name, isBot, Profile.Email, Profile.DisplayName, Profile.DisplayNameNormalized, Profile.RealName, Profile.RealNameNormalized, Profile.Image192, Profile.Image512
		_, _ = c.Writer.WriteString("ID: " + user.ID + "\n")
		_, _ = c.Writer.WriteString("Name: " + user.Name + "\n")
		_, _ = c.Writer.WriteString("isBot: " + fmt.Sprintf("%v", user.IsBot) + "\n")
		_, _ = c.Writer.WriteString("Profile.Email: " + user.Profile.Email + "\n")
		_, _ = c.Writer.WriteString("Profile.DisplayName: " + user.Profile.DisplayName + "\n")
		_, _ = c.Writer.WriteString("Profile.DisplayNameNormalized: " + user.Profile.DisplayNameNormalized + "\n")
		_, _ = c.Writer.WriteString("Profile.RealName: " + user.Profile.RealName + "\n")
		_, _ = c.Writer.WriteString("Profile.RealNameNormalized: " + user.Profile.RealNameNormalized + "\n")
		_, _ = c.Writer.WriteString("Profile.Image192: " + user.Profile.Image192 + "\n")
		_, _ = c.Writer.WriteString("Profile.Image512: " + user.Profile.Image512 + "\n")
		_, _ = c.Writer.WriteString("\n")
		result = append(result, user.ID)
	}
	_, _ = c.Writer.WriteString(fmt.Sprintf("%+q", result))
	_, _ = c.Writer.WriteString("\n")
	return result
}

// CAUTION
// This method applies only to direct message channels
// refer to https://api.slack.com/methods/conversations.close
func testCloseConversation(token string, channelID string, c *gin.Context) {
	if channelID == "" {
		return
	}
	noOp, alreadyClosed, err := api.CloseConversation(token, channelID)

	//print result to c
	_, _ = c.Writer.WriteString("---CloseConversation---" + "\n")
	_, _ = c.Writer.WriteString("noOp: " + fmt.Sprintf("%v", noOp) + "\n")
	_, _ = c.Writer.WriteString("alreadyClosed: " + fmt.Sprintf("%v", alreadyClosed) + "\n")
	_, _ = c.Writer.WriteString("err: " + fmt.Sprintf("%v", err) + "\n")

	if err != nil {
		_, _ = c.Writer.WriteString("Error in CloseConversation: " + err.Error() + "\n")
		return
	}
}

func testSendNotificationTemplate(token string, channelID string, mentionUser []string, c *gin.Context) {
	if channelID == "" {
		return
	}
	_, _ = c.Writer.WriteString("---SendNotificationTemplate---" + "\n")
	r1, r2, err := uc.SendRecordCompleteNotification(&api, mentionUser, channelID, token)
	if err != nil {
		_, _ = c.Writer.WriteString("Error in SendNotificationTemplate: " + err.Error() + "\n")
		return
	}
	_, _ = c.Writer.WriteString("ConversationID: " + r1 + "\n")
	_, _ = c.Writer.WriteString("TimestampID: " + r2 + "\n")

	_, _ = c.Writer.WriteString("---SendNotificationTemplate1---" + "\n")
	r3, r4, err := uc.SendRecordCompleteNotification1(&api, mentionUser, channelID, token)
	if err != nil {
		_, _ = c.Writer.WriteString("Error in SendNotificationTemplate1: " + err.Error() + "\n")
		return
	}
	_, _ = c.Writer.WriteString("ConversationID: " + r3 + "\n")
	_, _ = c.Writer.WriteString("TimestampID: " + r4 + "\n")
}

//func testGetUserIdentity(token string, c *gin.Context) *slack.UserIdentityResponse {
//	identity, err := api.GetUserIdentity(token)
//	if err != nil {
//		_, _ = c.Writer.WriteString("Error in GetUserIdentity: " + err.Error() + "\n")
//		return nil
//	}
//	_, _ = c.Writer.WriteString("---GetUserIdentity---" + "\n")
//
//	//Print all fields
//	_, _ = c.Writer.WriteString(fmt.Sprintf("%+v", identity))
//	return identity
//}

func testGetUserProfile(token string, c *gin.Context) *slack.UserProfile {
	profile, err := api.GetUserProfile(token)
	if err != nil {
		_, _ = c.Writer.WriteString("Error in GetUserProfile: " + err.Error() + "\n")
		return nil
	}
	_, _ = c.Writer.WriteString("---GetUserProfile---" + "\n")

	//Print all fields
	_, _ = c.Writer.WriteString(fmt.Sprintf("%+v", profile))
	return profile
}

func testGetUserInfo(token string, userID string, c *gin.Context) *slack.User {
	user, err := api.GetUserInfo(token, userID)
	if err != nil {
		_, _ = c.Writer.WriteString("Error in GetUserInfo: " + err.Error() + "\n")
		return nil
	}
	_, _ = c.Writer.WriteString("---GetUserInfo---" + "\n")

	//Print all fields
	_, _ = c.Writer.WriteString(fmt.Sprintf("%+v", user))
	return user
}

func testSendDirectMessage(token string, userID string, c *gin.Context) {
	if userID == "" {
		return
	}
	_, _ = c.Writer.WriteString("---SendDirectMessage---" + "\n")
	result, err := api.SendDirectMessage(token, userID, "Hello, World!")
	if err != nil {
		_, _ = c.Writer.WriteString("Error in SendDirectMessage: " + err.Error() + "\n")
		return
	} else {
		//print result to c
		_, _ = c.Writer.WriteString("Conversation ID: " + result.ConversationID + "\n")
		_, _ = c.Writer.WriteString("Timestamp ID: " + result.TimestampID + "\n")
	}
}

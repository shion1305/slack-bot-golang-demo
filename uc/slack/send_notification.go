package slack

import (
	slackLib "github.com/slack-go/slack"
	"goSlackBotEmpath/pkg/slack"
	"strings"
)

func createMentionMessage(mentionTargets []string) string {
	for _, user := range mentionTargets {
		user = "<@" + user + ">"
	}
	return strings.Join(mentionTargets, " ")
}

func SendRecordCompleteNotification(api *slack.SlackAPI, targetUsers []string, channelID string, token string) (conversationID string, timestampID string, err error) {
	attachment := slackLib.Attachment{
		Pretext:   "録画が完了しました!" + createMentionMessage(targetUsers),
		Color:     "#36a64f",
		Title:     "(MTG名)",
		TitleLink: "https://jam-roll.jp",
		Text:      "MTG時間: 2022/11/07\\nMTG 終了: 17:00",
		Fields: []slackLib.AttachmentField{
			{
				Title: "Field Title",
				Value: "Field Value",
				Short: false,
			},
			{
				Title: "Field Title",
				Value: "Field Value",
				Short: false,
			},
			{
				Title: "Field Title",
				Value: "Field Value",
				Short: false,
			},
			{
				Title: "Field Title",
				Value: "Field Value",
				Short: false,
			},
		},
	}
	return api.PostMessageWithAttachment(token, channelID, attachment)
}

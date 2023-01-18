package slack

import (
	slackLib "github.com/slack-go/slack"
	"goSlackBotEmpath/pkg/slack"
	"strings"
)

func CreateMentionMessage(mentionTargets []string) string {
	var targets []string
	for _, user := range mentionTargets {
		targets = append(targets, "<@"+user+">")
	}
	return strings.Join(targets, " ")
}

func SendRecordCompleteNotification(api *slack.SlackAPI, targetUsers []string, channelID string, token string) (conversationID string, timestampID string, err error) {
	attachment := slackLib.Attachment{
		Pretext:   "録画が完了しました!" + CreateMentionMessage(targetUsers),
		Color:     "#36a64f",
		Title:     "(MTG名)",
		TitleLink: "https://jam-roll.jp",
		Fields: []slackLib.AttachmentField{
			{Title: "日時", Value: "2022/11/07", Short: false},
			{Title: "MTG開始", Value: "12:14", Short: true},
			{Title: "MTG終了", Value: "15:30", Short: true},
		},
		AuthorName: "Author",
		AuthorIcon: "https://img.freepik.com/premium-vector/cartoon-cute-square-shaped-frog-square-icon-apps-games-vector-illustration-isolated_351178-27.jpg",
	}
	return api.PostMessageWithAttachment(token, channelID, attachment)
}

func SendRecordCompleteNotification1(api *slack.SlackAPI, targetUsers []string, channelID string, token string) (conversationID string, timestampID string, err error) {
	attachment := slackLib.Attachment{
		Pretext:   CreateMentionMessage(targetUsers),
		Color:     "#36a64f",
		Title:     "録画が完了しました!",
		TitleLink: "https://jam-roll.jp",
		Fields: []slackLib.AttachmentField{
			{Title: "MTG名", Value: "テストミーティング", Short: true},
			{Title: "日時", Value: "2022/11/07", Short: true},
			{Title: "MTG開始", Value: "12:14", Short: true},
			{Title: "MTG終了", Value: "15:30", Short: true},
		},
		AuthorName: "Author",
		AuthorIcon: "https://img.freepik.com/premium-vector/cartoon-cute-square-shaped-frog-square-icon-apps-games-vector-illustration-isolated_351178-27.jpg",
	}
	return api.PostMessageWithAttachment(token, channelID, attachment)
}

func SendChannelChangeNotification(api *slack.SlackAPI, channelID string, token string) (conversationID string, timestampID string, err error) {
	attachment := slackLib.Attachment{
		Title:      "通知チャンネルが変更されました!",
		Color:      "#36a64f",
		AuthorName: "Author",
		AuthorIcon: "https://img.freepik.com/premium-vector/cartoon-cute-square-shaped-frog-square-icon-apps-games-vector-illustration-isolated_351178-27.jpg",
		Text:       "設定が変更されたため、通知チャンネルが変更されました。",
	}
	return api.PostMessageWithAttachment(token, channelID, attachment)
}

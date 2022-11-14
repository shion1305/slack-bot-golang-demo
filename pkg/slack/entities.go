package slack

import "encoding/json"

type Channel struct {
	GroupConversation
	IsChannel bool   `json:"is_channel"`
	IsGeneral bool   `json:"is_general"`
	IsMember  bool   `json:"is_member"`
	Locale    string `json:"locale"`
}

type GroupConversation struct {
	Conversation
	Name       string   `json:"name"`
	Creator    string   `json:"creator"`
	IsArchived bool     `json:"is_archived"`
	Members    []string `json:"members"`
	Topic      Topic    `json:"topic"`
	Purpose    Purpose  `json:"purpose"`
}

type Conversation struct {
	ID                 string   `json:"id"`
	Created            JSONTime `json:"created"`
	IsOpen             bool     `json:"is_open"`
	LastRead           string   `json:"last_read,omitempty"`
	Latest             *Message `json:"latest,omitempty"`
	UnreadCount        int      `json:"unread_count,omitempty"`
	UnreadCountDisplay int      `json:"unread_count_display,omitempty"`
	IsGroup            bool     `json:"is_group"`
	IsShared           bool     `json:"is_shared"`
	IsIM               bool     `json:"is_im"`
	IsExtShared        bool     `json:"is_ext_shared"`
	IsOrgShared        bool     `json:"is_org_shared"`
	IsPendingExtShared bool     `json:"is_pending_ext_shared"`
	IsPrivate          bool     `json:"is_private"`
	IsMpIM             bool     `json:"is_mpim"`
	Unlinked           int      `json:"unlinked"`
	NameNormalized     string   `json:"name_normalized"`
	NumMembers         int      `json:"num_members"`
	Priority           float64  `json:"priority"`
	User               string   `json:"user"`
}

type Topic struct {
	Value   string   `json:"value"`
	Creator string   `json:"creator"`
	LastSet JSONTime `json:"last_set"`
}

type Purpose struct {
	Value   string   `json:"value"`
	Creator string   `json:"creator"`
	LastSet JSONTime `json:"last_set"`
}

type Message struct {
	Msg
	SubMessage      *Msg `json:"message,omitempty"`
	PreviousMessage *Msg `json:"previous_message,omitempty"`
}

type Msg struct {
	// Basic Message
	ClientMsgID     string       `json:"client_msg_id,omitempty"`
	Type            string       `json:"type,omitempty"`
	Channel         string       `json:"channel,omitempty"`
	User            string       `json:"user,omitempty"`
	Text            string       `json:"text,omitempty"`
	Timestamp       string       `json:"ts,omitempty"`
	ThreadTimestamp string       `json:"thread_ts,omitempty"`
	IsStarred       bool         `json:"is_starred,omitempty"`
	PinnedTo        []string     `json:"pinned_to,omitempty"`
	Attachments     []Attachment `json:"attachments,omitempty"`
	LastRead        string       `json:"last_read,omitempty"`
	Subscribed      bool         `json:"subscribed,omitempty"`
	UnreadCount     int          `json:"unread_count,omitempty"`

	// Message Subtypes
	SubType string `json:"subtype,omitempty"`

	// Hidden Subtypes
	Hidden           bool   `json:"hidden,omitempty"`     // message_changed, message_deleted, unpinned_item
	DeletedTimestamp string `json:"deleted_ts,omitempty"` // message_deleted
	EventTimestamp   string `json:"event_ts,omitempty"`

	// bot_message (https://api.slack.com/events/message/bot_message)
	BotID    string `json:"bot_id,omitempty"`
	Username string `json:"username,omitempty"`

	// channel_join, group_join
	Inviter string `json:"inviter,omitempty"`

	// channel_topic, group_topic
	Topic string `json:"topic,omitempty"`

	// channel_purpose, group_purpose
	Purpose string `json:"purpose,omitempty"`

	// channel_name, group_name
	Name    string `json:"name,omitempty"`
	OldName string `json:"old_name,omitempty"`

	// channel_archive, group_archive
	Members []string `json:"members,omitempty"`

	// channels.replies, groups.replies, im.replies, mpim.replies
	ReplyCount   int    `json:"reply_count,omitempty"`
	ParentUserId string `json:"parent_user_id,omitempty"`
	LatestReply  string `json:"latest_reply,omitempty"`

	// file_share
	Upload bool `json:"upload,omitempty"`

	// pinned_item
	ItemType string `json:"item_type,omitempty"`

	// https://api.slack.com/rtm
	ReplyTo int    `json:"reply_to,omitempty"`
	Team    string `json:"team,omitempty"`

	// slash commands and interactive messages
	ResponseType    string `json:"response_type,omitempty"`
	ReplaceOriginal bool   `json:"replace_original"`
	DeleteOriginal  bool   `json:"delete_original"`

	// permalink
	Permalink string `json:"permalink,omitempty"`
}

type Attachment struct {
	Color    string `json:"color,omitempty"`
	Fallback string `json:"fallback,omitempty"`

	CallbackID string `json:"callback_id,omitempty"`
	ID         int    `json:"id,omitempty"`

	AuthorID      string `json:"author_id,omitempty"`
	AuthorName    string `json:"author_name,omitempty"`
	AuthorSubname string `json:"author_subname,omitempty"`
	AuthorLink    string `json:"author_link,omitempty"`
	AuthorIcon    string `json:"author_icon,omitempty"`

	Title     string `json:"title,omitempty"`
	TitleLink string `json:"title_link,omitempty"`
	Pretext   string `json:"pretext,omitempty"`
	Text      string `json:"text,omitempty"`

	ImageURL string `json:"image_url,omitempty"`
	ThumbURL string `json:"thumb_url,omitempty"`

	ServiceName string `json:"service_name,omitempty"`
	ServiceIcon string `json:"service_icon,omitempty"`
	FromURL     string `json:"from_url,omitempty"`
	OriginalURL string `json:"original_url,omitempty"`

	MarkdownIn []string `json:"mrkdwn_in,omitempty"`

	Footer     string `json:"footer,omitempty"`
	FooterIcon string `json:"footer_icon,omitempty"`

	Ts json.Number `json:"ts,omitempty"`
}

type Team struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Enterprise struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type AuthedUser struct {
	ID           string `json:"id"`
	Scope        string `json:"scope"`
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
}

package slack

//
//import (
//	"fmt"
//	"time"
//)
//
//// AccessResponse struct for reading oauth.v2.access response
//type AccessResponse struct {
//	AccessToken         string     `json:"access_token"`
//	TokenType           string     `json:"token_type"`
//	Scope               string     `json:"scope"`
//	BotUserID           string     `json:"bot_user_id"`
//	AppID               string     `json:"app_id"`
//	ExpiresIn           int        `json:"expires_in"`
//	RefreshToken        string     `json:"refresh_token"`
//	Team                Team       `json:"team"`
//	Enterprise          Enterprise `json:"enterprise"`
//	AuthedUser          AuthedUser `json:"authed_user"`
//	IsEnterpriseInstall bool       `json:"is_enterprise_install"`
//	ResponseStatus
//}
//
//type ConversationListResponse struct {
//	Channels         []Channel        `json:"channels"`
//	ResponseMetaData ResponseMetadata `json:"response_metadata"`
//	ResponseStatus
//}
//
//type ConversationCreateResponse struct {
//	Channel Channel `json:"channel"`
//	ResponseStatus
//}
//
//type ResponseStatus struct {
//	Ok               bool             `json:"ok"`
//	Error            bool             `json:"error"`
//	ResponseMetadata ResponseMetadata `json:"response_metadata"`
//}
//
//type ResponseMetadata struct {
//	Cursor   string   `json:"next_cursor"`
//	Messages []string `json:"messages"`
//	Warnings []string `json:"warnings"`
//}
//
//// JSONTime exists so that we can have a String method converting the date
//type JSONTime int64
//
//// String converts the unix timestamp into a string
//func (t JSONTime) String() string {
//	tm := t.Time()
//	return fmt.Sprintf("\"%s\"", tm.Format("Mon Jan _2"))
//}
//
//// Time returns a `time.Time` representation of this value.
//func (t JSONTime) Time() time.Time {
//	return time.Unix(int64(t), 0)
//}

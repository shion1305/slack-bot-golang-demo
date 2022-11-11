package slack

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

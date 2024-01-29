package user

type User struct {
	Username   string   `json:"username"`
	Password   string   `json:"password"`
	AuthTokens []string `json:"authTokens"`
	BotTokens  []string `json:"botTokens"`
}

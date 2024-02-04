package user

type User struct {
	Email      string   `json:"email"`
	Password   string   `json:"password"`
	AuthTokens []string `json:"authTokens"`
	BotTokens  []string `json:"botTokens"`
}

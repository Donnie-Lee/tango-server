package loginModel

type Token struct {
	TokenName  string `json:"tokenName"`
	TokenValue string `json:"tokenValue"`
	LoginId    int    `json:"loginId"`
}

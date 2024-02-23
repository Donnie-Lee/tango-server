package loginModel

type LoginSmsRequest struct {
	Mobile    string `json:"mobile"`
	CheckCode string `json:"checkcode"`
}

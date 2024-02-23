package accountModel

type ContactInfoResponse struct {
	AccountInfo    *AccountExt `json:"accountInfo"`
	CreateTime     string      `json:"createTime"`
	Status         int         `json:"status"`
	Stared         bool        `json:"stared"`
	NickNameRemark string      `json:"nickNameRemark"`
}

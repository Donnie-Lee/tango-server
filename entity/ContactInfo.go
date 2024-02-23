package entity

import "time"

type ContactInfo struct {
	ID             int        `gorm:"primaryKey" json:"id"`
	UserId         int        `json:"userId"`
	ContactUserId  int        `json:"contactUserId"`
	CreateTime     *time.Time `gorm:"autoCreateTime" json:"createTime"`
	Stared         MyBool     `json:"stared"`
	NickNameRemark string     `json:"nickNameRemark"`
}

func (ContactInfo) TableName() string {
	return "contact_info"
}

package entity

import "time"

type AccountInfo struct {
	ID         int        `gorm:"primaryKey" json:"id"`
	Nickname   string     `json:"nickname"`
	Mobile     string     `json:"mobile"`
	Avatar     string     `json:"avatar"`
	BgImage    string     `json:"bgImage"`
	Introduce  string     `json:"introduce"`
	SignInfo   string     `json:"signInfo"`
	Gender     string     `json:"gender"`
	CreateTime *time.Time `gorm:"autoUpdateTime" json:"createTime"`
	UpdateTime *time.Time `gorm:"autoCreateTime" json:"updateTime"`
	Revision   uint64     `json:"revision"`
	Password   string     `json:"password"`
}

func (AccountInfo) TableName() string {
	return "account_info"
}

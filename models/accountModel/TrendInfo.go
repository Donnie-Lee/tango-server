package accountModel

import "time"

type TrendInfo struct {
	Id         int64      `json:"id"`
	Images     []string   `json:"images"`
	Content    string     `json:"content"`
	CreateTime *time.Time `json:"createTime"`
}

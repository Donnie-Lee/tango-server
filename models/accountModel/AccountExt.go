package accountModel

import "imserver/entity"

type AccountExt struct {
	*entity.AccountInfo
	CreateTime string      `json:"createTime"`
	UpdateTime string      `json:"updateTime"`
	Trends     []TrendInfo `json:"trends"`
}

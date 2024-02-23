package account

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"imserver/common/jwt"
	"imserver/common/response"
	"imserver/entity"
	"imserver/middleware/database"
	"imserver/models/accountModel"
	"imserver/models/loginModel"
)

func LoginSms(mobile string, ctx *gin.Context) {
	var account *entity.AccountInfo
	result := database.DB.Where("mobile =? ", mobile).First(&account)
	if result.RowsAffected == 0 {
		account = &entity.AccountInfo{
			Nickname: mobile,
			Mobile:   mobile,
			Gender:   "SECRET",
		}
		result = database.DB.Create(&account)
		if result.Error != nil {
			panic(result.Error)
		}
	}

	tokenStr := jwt.GetToken(account.ID, account.Mobile)

	ctx.JSON(200, response.SuccessWithData(loginModel.Token{
		TokenName:  "Authorization",
		TokenValue: tokenStr,
		LoginId:    account.ID,
	}))
}

func GetAccountInfo(id int, ctx *gin.Context) {
	var account *entity.AccountInfo
	database.DB.First(&account, id)
	accountExt := transferToAccountExt(account)
	accountExt.Trends = []accountModel.TrendInfo{}
	accountExt.CreateTime = account.CreateTime.Format("2006-01-02 15:04:05")
	accountExt.UpdateTime = account.UpdateTime.Format("2006-01-02 15:04:05")
	ctx.JSON(200, response.SuccessWithData(accountExt))

}

func transferToAccountExt(account *entity.AccountInfo) *accountModel.AccountExt {
	var accountExt *accountModel.AccountExt
	temp, _ := json.Marshal(account)
	json.Unmarshal(temp, &accountExt)
	accountExt.Trends = []accountModel.TrendInfo{}
	return accountExt
}

func Contacts(ctx *gin.Context) {
	claims, _ := ctx.Get(`claims`)
	var contactInfos []*entity.ContactInfo
	result := database.DB.Where("user_id = ?", claims.(jwt.UserClaims).LoginId).Find(&contactInfos)

	if result.RowsAffected == 0 {
		ctx.JSON(200, response.SuccessWithData(contactInfos))
		return
	}
	var ids []int
	for _, info := range contactInfos {
		ids = append(ids, info.ContactUserId)
	}
	var accountInfos []*entity.AccountInfo
	database.DB.Find(&accountInfos, ids)

	var contactInfoResponses []*accountModel.ContactInfoResponse
	for _, contactInfo := range contactInfos {
		contactInfoResponse := transferToContactInfoResponses(contactInfo)
		for _, info := range accountInfos {
			if info.ID == contactInfo.ContactUserId {
				contactInfoResponse.AccountInfo = transferToAccountExt(info)
			}
		}
		contactInfoResponses = append(contactInfoResponses, contactInfoResponse)
	}

	ctx.JSON(200, response.SuccessWithData(contactInfoResponses))

}

func transferToContactInfoResponses(contactInfo *entity.ContactInfo) *accountModel.ContactInfoResponse {
	var contactInfoResponse *accountModel.ContactInfoResponse
	temp, _ := json.Marshal(contactInfo)
	json.Unmarshal(temp, &contactInfoResponse)
	return contactInfoResponse
}

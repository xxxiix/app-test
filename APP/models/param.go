package models

// 定义请求的参数结构体
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
	Phone      string `json:"phone" binding:"required"`
}

type ParamLogin struct {
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ParamSignUpCheckInfo struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
	Phone      string `json:"phone" binding:"required"`
	Code       string `json:"code" binding:"required"`
}

type ParamAddBill struct {
	User_id           int64
	Income_or_outcome string `json:"income_or_outcome" binding:"required"`
	Money             int64  `json:"money" binding:"required"`
	Bill_type         string `json:"bill_type" binding:"required"`
	Bill_info         string `json:"bill_info"`
	Bill_day          string `json:"bill_day" binding:"required"`
}

type ParamDeleteBill struct {
	Bill_id int64 `json:"bill_id" bind:"required"`
	User_id int64
}

type ParamChangeBill struct {
	Bill_id           int64 `json:"bill_id" bind:"required"`
	User_id           int64
	Income_or_outcome string `json:"income_or_outcome" binding:"required"`
	Money             int64  `json:"money" binding:"required"`
	Bill_type         string `json:"bill_type" binding:"required"`
	Bill_info         string `json:"bill_info"`
	Bill_day          string `json:"bill_day" binding:"required"`
}

type ParamSearchBillByDay struct {
	Bill_day          string `json:"bill_day" bind:"required"`
	User_id           int64
	Income_or_outcome string `json:"income_or_outcome"`
}

type ParamSearchBillByWeek struct {
	Bill_isoweek      string `json:"bill_isoweek" bind:"required"`
	User_id           int64
	Income_or_outcome string `json:"income_or_outcome"`
}

type ParamSearchBillByMonth struct {
	Bill_month        string `json:"bill_month" bind:"required"`
	User_id           int64
	Income_or_outcome string `json:"income_or_outcome"`
}

type ParamSearchBillByYear struct {
	Bill_year         string `json:"bill_year" bind:"required"`
	User_id           int64
	Income_or_outcome string `json:"income_or_outcome"`
}

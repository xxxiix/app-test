package controllers

import (
	"errors"
	"fmt"
	"main/dao/mysql"
	"main/logic"
	"main/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func AddBillHandler(c *gin.Context) {
	// 参数校验
	userID, err := CheckUser(c)
	if err != nil {
		zap.L().Error("User error", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	p := new(models.ParamAddBill)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Add bill with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	if p.Income_or_outcome != "income" && p.Income_or_outcome != "outcome" {
		ResponseError(c, CodeInvalidParam)
		return
	}
	if p.Money <= 0 {
		ResponseError(c, CodeInvalidParam)
		return
	}
	if _, err := time.Parse(logic.TimeFormAll, p.Bill_day); err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	p.User_id = userID

	// 业务逻辑处理
	if err := logic.AddBill(p); err != nil {
		zap.L().Error("Keep data failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil, nil)
	fmt.Println(p)
}

// 留给账单照片上传用的
func AddByPhotoBillHandler(c *gin.Context) {

}

func DeleteBillHandler(c *gin.Context) {
	// 参数校验
	userID, err := CheckUser(c)
	if err != nil {
		zap.L().Error("User error", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	p := new(models.ParamDeleteBill)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Delete bill with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	p.User_id = userID

	// 业务逻辑处理
	if err := logic.DeleteBill(p); err != nil {
		zap.L().Error("Delete data failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil, nil)
	fmt.Println(p)
}

func ChangeBillHandler(c *gin.Context) {
	// 参数校验
	userID, err := CheckUser(c)
	if err != nil {
		zap.L().Error("User error", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	p := new(models.ParamChangeBill)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Change bill with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	if p.Income_or_outcome != "income" && p.Income_or_outcome != "outcome" {
		ResponseError(c, CodeInvalidParam)
		return
	}
	if p.Money <= 0 {
		ResponseError(c, CodeInvalidParam)
		return
	}
	if _, err := time.Parse(logic.TimeFormAll, p.Bill_day); err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	p.User_id = userID

	// 业务逻辑处理
	if err := logic.ChangeBill(p); err != nil {
		zap.L().Error("Change data failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil, nil)
	fmt.Println(p)
}

func SearchBillHandlerByDay(c *gin.Context) {
	// 参数校验
	userID, err := CheckUser(c)
	if err != nil {
		zap.L().Error("User error", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	p := new(models.ParamSearchBillByDay)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Search bill by day with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	p.User_id = userID

	// 业务逻辑处理
	data, err := logic.SearchBillByDay(p)
	if err != nil {
		zap.L().Error("Search data by day failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorBillNotExist) {
			ResponseError(c, CodeBillNotExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data, nil)
	fmt.Println(p)
}

func SearchBillHandlerByWeek(c *gin.Context) {
	// 参数校验
	userID, err := CheckUser(c)
	if err != nil {
		zap.L().Error("User error", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	p := new(models.ParamSearchBillByWeek)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Search bill by week with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	p.User_id = userID

	// 业务逻辑处理
	data, err := logic.SearchBillByWeek(p)
	if err != nil {
		zap.L().Error("Search data by week failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorBillNotExist) {
			ResponseError(c, CodeBillNotExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data, nil)
	fmt.Println(p)
}

func SearchBillHandlerByMonth(c *gin.Context) {
	// 参数校验
	userID, err := CheckUser(c)
	if err != nil {
		zap.L().Error("User error", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	p := new(models.ParamSearchBillByMonth)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Search bill by month with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	p.User_id = userID

	// 业务逻辑处理
	data, err := logic.SearchBillByMonth(p)
	if err != nil {
		zap.L().Error("Search data by month failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorBillNotExist) {
			ResponseError(c, CodeBillNotExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data, nil)
	fmt.Println(p)
}

func SearchBillHandlerByYear(c *gin.Context) {
	// 参数校验
	userID, err := CheckUser(c)
	if err != nil {
		zap.L().Error("User error", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	p := new(models.ParamSearchBillByYear)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Search bill by year with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	p.User_id = userID

	// 业务逻辑处理
	data, err := logic.SearchBillByYear(p)
	if err != nil {
		zap.L().Error("Search data by year failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorBillNotExist) {
			ResponseError(c, CodeBillNotExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data, nil)
	fmt.Println(p)
}

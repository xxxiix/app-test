package logic

import (
	"fmt"
	"main/dao/mysql"
	"main/models"
	"main/pkg/snowflake"
	"time"
)

const (
	TimeFormAll   = "2006-01-02"
	TimeFormYear  = "2006"
	TimeFormMomth = "2006-01"
)

func AddBill(p *models.ParamAddBill) (err error) {
	b := new(models.Bill)
	b.User_id = p.User_id
	b.Bill_id = snowflake.GenID()
	b.Income_or_outcome = p.Income_or_outcome
	b.Money = p.Money
	b.Bill_type = p.Bill_type
	b.Bill_info = p.Bill_info
	b.Bill_date = p.Bill_day
	t, _ := time.Parse(TimeFormAll, p.Bill_day)
	b.Bill_year = t.Format(TimeFormYear)
	b.Bill_month = t.Format(TimeFormMomth)
	y, w := t.ISOWeek()
	b.Bill_isoweek = fmt.Sprintf("%d-%d", y, w)

	return mysql.InsertBill(b)
}

func DeleteBill(p *models.ParamDeleteBill) (err error) {
	b := new(models.Bill)
	b.User_id = p.User_id
	b.Bill_id = p.Bill_id

	return mysql.DeleteById(b)
}

func ChangeBill(p *models.ParamChangeBill) (err error) {
	b := new(models.Bill)
	b.User_id = p.User_id
	b.Bill_id = p.Bill_id
	b.Income_or_outcome = p.Income_or_outcome
	b.Money = p.Money
	b.Bill_type = p.Bill_type
	b.Bill_info = p.Bill_info
	b.Bill_date = p.Bill_day
	t, _ := time.Parse(TimeFormAll, p.Bill_day)
	b.Bill_year = t.Format(TimeFormYear)
	b.Bill_month = t.Format(TimeFormMomth)
	y, w := t.ISOWeek()
	b.Bill_isoweek = fmt.Sprintf("%d.%d", y, w)

	return mysql.ChangeById(b)
}

func SearchBillByDay(p *models.ParamSearchBillByDay) (back *[]models.BackBill, err error) {
	b := new(models.Bill)
	b.User_id = p.User_id
	b.Bill_date = p.Bill_day
	if p.Income_or_outcome != "" {
		b.Income_or_outcome = p.Income_or_outcome
	}

	return mysql.SearchByDay(b)
}

func SearchBillByWeek(p *models.ParamSearchBillByWeek) (back *[]models.BackBill, err error) {
	b := new(models.Bill)
	b.User_id = p.User_id
	b.Bill_isoweek = p.Bill_isoweek
	if p.Income_or_outcome != "" {
		b.Income_or_outcome = p.Income_or_outcome
	}

	return mysql.SearchByWeek(b)
}

func SearchBillByMonth(p *models.ParamSearchBillByMonth) (back *[]models.BackBill, err error) {
	b := new(models.Bill)
	b.User_id = p.User_id
	b.Bill_month = p.Bill_month
	if p.Income_or_outcome != "" {
		b.Income_or_outcome = p.Income_or_outcome
	}

	return mysql.SearchByMonth(b)
}

func SearchBillByYear(p *models.ParamSearchBillByYear) (back *[]models.BackBill, err error) {
	b := new(models.Bill)
	b.User_id = p.User_id
	b.Bill_year = p.Bill_year
	if p.Income_or_outcome != "" {
		b.Income_or_outcome = p.Income_or_outcome
	}

	return mysql.SearchByYear(b)
}

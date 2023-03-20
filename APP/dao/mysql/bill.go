package mysql

import (
	"database/sql"
	"main/models"
)

func InsertBill(b *models.Bill) (err error) {
	sqlStr := `insert into 
	bill(user_id, bill_id, income_or_outcome, money, bill_type,
		bill_info, bill_year, bill_month, bill_isoweek, bill_date)
	value(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err = db.Exec(sqlStr,
		b.User_id, b.Bill_id, b.Income_or_outcome, b.Money,
		b.Bill_type, b.Bill_info, b.Bill_year, b.Bill_month,
		b.Bill_isoweek, b.Bill_date)
	return
}

// 有Bill_id的情况下，可以和User_id进行互相检验，但是没有就不能用这个
func CheckBillIsUser(b *models.Bill) (err error) {
	userID := b.User_id
	sqlStr := `select user_id from bill where bill_id=?`
	if err = db.Get(b, sqlStr, b.Bill_id); err != nil {
		if err == sql.ErrNoRows {
			return ErrorInvalidID
		}
		return
	}
	if userID != b.User_id {
		return ErrorInvalidID
	}
	return
}

func ChangeById(b *models.Bill) (err error) {
	if err = CheckBillIsUser(b); err != nil {
		return
	}
	sqlStr := `update bill set
	income_or_outcome=?, money=?,
	bill_type=?, bill_info=?, bill_year=?, 
	bill_month=?, bill_isoweek=?,bill_date=?
	where bill_id=?`
	_, err = db.Exec(sqlStr,
		b.Income_or_outcome, b.Money,
		b.Bill_type, b.Bill_info, b.Bill_year,
		b.Bill_month, b.Bill_isoweek, b.Bill_date,
		b.Bill_id)
	return
}

func DeleteById(b *models.Bill) (err error) {
	if err = CheckBillIsUser(b); err != nil {
		return
	}
	sqlStr := `delete 
	from bill 
	where bill_id=?`
	_, err = db.Exec(sqlStr, b.Bill_id)
	return
}

func SearchByDay(b *models.Bill) (back *[]models.BackBill, err error) {
	var sqlStr string
	back = new([]models.BackBill)
	if b.Income_or_outcome != "" {
		sqlStr = `select
		bill_id, income_or_outcome, money,
		bill_type, bill_info, bill_date, bill_isoweek
		from bill
		where user_id=? and bill_date=? and income_or_outcome=?`
		if err = db.Select(back, sqlStr, b.User_id, b.Bill_date, b.Income_or_outcome); err != nil {
			if err == sql.ErrNoRows {
				err = ErrorBillNotExist
				return
			}
			return
		}
	}
	sqlStr = `select
		bill_id, income_or_outcome, money,
		bill_type, bill_info, bill_date, bill_isoweek
		from bill
		where user_id=? and bill_date=?`
	if err = db.Select(back, sqlStr, b.User_id, b.Bill_date); err != nil {
		if err == sql.ErrNoRows {
			err = ErrorBillNotExist
			return
		}
		return
	}
	return
}

func SearchByWeek(b *models.Bill) (back *[]models.BackBill, err error) {
	var sqlStr string
	back = new([]models.BackBill)
	if b.Income_or_outcome != "" {
		sqlStr = `select
		bill_id, income_or_outcome, money,
		bill_type, bill_info, bill_date, bill_isoweek
		from bill
		where user_id=? and bill_isoweek=? and income_or_outcome=?`
		if err = db.Select(back, sqlStr, b.User_id, b.Bill_isoweek, b.Income_or_outcome); err != nil {
			if err == sql.ErrNoRows {
				err = ErrorBillNotExist
				return
			}
			return
		}
	}
	sqlStr = `select
		bill_id, income_or_outcome, money,
		bill_type, bill_info, bill_date, bill_isoweek
		from bill
		where user_id=? and bill_isoweek=?`
	if err = db.Select(back, sqlStr, b.User_id, b.Bill_isoweek); err != nil {
		if err == sql.ErrNoRows {
			err = ErrorBillNotExist
			return
		}
		return
	}
	return
}

func SearchByMonth(b *models.Bill) (back *[]models.BackBill, err error) {
	var sqlStr string
	back = new([]models.BackBill)
	if b.Income_or_outcome != "" {
		sqlStr = `select
		bill_id, income_or_outcome, money,
		bill_type, bill_info, bill_date, bill_isoweek
		from bill
		where user_id=? and bill_month=? and income_or_outcome=?`
		if err = db.Select(back, sqlStr, b.User_id, b.Bill_month, b.Income_or_outcome); err != nil {
			if err == sql.ErrNoRows {
				err = ErrorBillNotExist
				return
			}
			return
		}
	}
	sqlStr = `select
		bill_id, income_or_outcome, money,
		bill_type, bill_info, bill_date, bill_isoweek
		from bill
		where user_id=? and bill_month=?`
	if err = db.Select(back, sqlStr, b.User_id, b.Bill_month); err != nil {
		if err == sql.ErrNoRows {
			err = ErrorBillNotExist
			return
		}
		return
	}
	return
}

func SearchByYear(b *models.Bill) (back *[]models.BackBill, err error) {
	var sqlStr string
	back = new([]models.BackBill)
	if b.Income_or_outcome != "" {
		sqlStr = `select
		bill_id, income_or_outcome, money,
		bill_type, bill_info, bill_date, bill_isoweek
		from bill
		where user_id=? and bill_year=? and income_or_outcome=?`
		if err = db.Select(back, sqlStr, b.User_id, b.Bill_year, b.Income_or_outcome); err != nil {
			if err == sql.ErrNoRows {
				err = ErrorBillNotExist
				return
			}
			return
		}
	}
	sqlStr = `select
		bill_id, income_or_outcome, money,
		bill_type, bill_info, bill_date, bill_isoweek
		from bill
		where user_id=? and bill_year=?`
	if err = db.Select(back, sqlStr, b.User_id, b.Bill_year); err != nil {
		if err == sql.ErrNoRows {
			err = ErrorBillNotExist
			return
		}
		return
	}
	return
}

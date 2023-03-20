package models

type Bill struct {
	User_id           int64  `db:"user_id"`
	Bill_id           int64  `db:"bill_id"`
	Income_or_outcome string `db:"income_or_outcome"`
	Money             int64  `db:"money"`
	Bill_type         string `db:"bill_type"`
	Bill_info         string `db:"bill_info"`
	Bill_year         string `db:"bill_year"`
	Bill_month        string `db:"bill_month"`
	Bill_isoweek      string `db:"bill_isoweek"`
	Bill_date         string `db:"bill_date"`
}

package models

type BackBill struct {
	Bill_id           int64  `json:"bill_id" db:"bill_id"`
	Income_or_outcome string `json:"income_or_outcome" db:"income_or_outcome"`
	Money             int64  `json:"money" db:"money"`
	Bill_type         string `json:"bill_type" db:"bill_type"`
	Bill_info         string `json:"bill_info" db:"bill_info"`
	Bill_isoweek      string `json:"bill_isoweek" db:"bill_isoweek"`
	Bill_day          string `json:"bill_day" db:"bill_date"`
}

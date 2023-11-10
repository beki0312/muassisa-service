package models

import "time"

type PosTerminalTransaction struct {
	OrderId        string `json:"orderId"`
	TransactionId  string `json:"transactionId"`
	TerminalNumber string `json:"terminalNumber"`
	Amount         string `json:"amount"`
	Status         string `json:"status"`
}

type AddTransaction struct {
	OrderId        string `json:"order_id" gorm:"column:order_id"`
	TransactionId  string `json:"transaction_id" gorm:"column:transaction_id"`
	TerminalNumber string `json:"terminal_number" gorm:"column:terminal_number"`
	Amount         string `json:"amount" gorm:"column:amount"`
	Status         string `json:"status" gorm:"column:status"`
	Type           string `json:"type" gorm:"column:type"`
	DateCreteCFT   string `json:"date_crete_cft" gorm:"column:create_at_cft"`
}
type AddedNewTransaction struct {
	OrderId        string `json:"order_id"`
	TransactionId  string `json:"transaction_id"`
	TerminalNumber string `json:"terminal_number"`
	Amount         string `json:"amount"`
	Status         string `json:"status"`
	Type           string `json:"type"`
	DateCreteCFT   string `json:"date_crete_cft"`
	Inn            string `json:"inn"`
}

type GetCourse struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Amount      string    `json:"amount"`
	DateBegin   time.Time `json:"date_begin"`
	DateEnd     time.Time `json:"date_end"`
	Created_At  time.Time `json:"created_At"`
}

type Language struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

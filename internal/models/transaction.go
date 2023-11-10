package models

import "time"

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

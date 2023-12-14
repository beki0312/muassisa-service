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
type GetCourseNew struct {
	Id                    int64     `json:"id"`
	CourseName            string    `json:"courseName" gorm:"column:name"`
	Title                 string    `json:"title"`
	CourseDescription     string    `json:"courseDescription" gorm:"column:description"`
	Amount                string    `json:"amount"`
	ImageName             string    `json:"image_name" gorm:"column:image_name"`
	Image                 string    `json:"image" gorm:"column:image"`
	Category              string    `json:"category"`
	Status                string    `json:"status"`
	InstructorName        string    `json:"instructorName" gorm:"column:name"`
	InsrtuctorUnvon       string    `json:"insrtuctorUnvon" gorm:"column:unvon"`
	InstructorDescription string    `json:"instructorDescription" gorm:"column:description"`
	Lectures              int       `json:"lectures"`
	Quizzes               int       `json:"quizzes"`
	Duration              string    `json:"duration"`
	SkillLevel            string    `json:"skill_Level"`
	Language              string    `json:"language"`
	Students              int       `json:"students"`
	Assessments           string    `json:"assessments"`
	DateBegin             time.Time `json:"date_begin"`
	DateEnd               time.Time `json:"date_end"`
	CreatedAt             time.Time `json:"created_At"`
}

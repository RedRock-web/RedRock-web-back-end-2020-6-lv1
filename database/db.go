package database

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var G_db *gorm.DB

type StudentRec struct {
	Code       int    `json:"code"`
	Info       string `json:"info"`
	ReturnData []struct {
		Xh    string      `json:"xh"`
		Xm    string      `json:"xm"`
		XmEn  interface{} `json:"xmEn"`
		Xb    string      `json:"xb"`
		Bj    string      `json:"bj"`
		Zyh   string      `json:"zyh"`
		Zym   string      `json:"zym"`
		Yxh   string      `json:"yxh"`
		Yxm   string      `json:"yxm"`
		Nj    string      `json:"nj"`
		Csrq  string      `json:"csrq"`
		Xjzt  string      `json:"xjzt"`
		Rxrq  string      `json:"rxrq"`
		Yxmen string      `json:"yxmen"`
		ZymEn string      `json:"zymEn"`
		Xz    int         `json:"xz"`
		Mz    string      `json:"mz"`
	} `json:"returnData"`
}

type Class struct {
	gorm.Model
	Name      string            `gorm:"index:idx_class_uix"`
	StudentId int               `gorm:"index:idx_class_uix"`
	ClassId   string
	Location  string
	Day       string             `gorm:"index:idx_class_uix"`
	Lesson    string             `gorm:"index:idx_class_uix"`
	RawWeek   string
	Teacher   string
	Semester  string
}

type Student struct {
	gorm.Model
	Grade      string
	Name       string
	StudentId  string    `gorm:"unique_index"`
	Gender     string
	CalssId    string
	Department string
	Profession string
	Status     string
	Nation     string
}

func Start()  {
	ConnetDb()
	CreateTable()
}

func ConnetDb() {
	db, err := gorm.Open("mysql", "root:mima@/students?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		errors.New("open database failed!")
	}
	G_db = db
}

func CreateTable() {
	if G_db.HasTable(&Student{}) {
		G_db.AutoMigrate()
	} else {
		G_db.CreateTable(&Student{})
	}
	if G_db.HasTable(&Class{}) {
		G_db.AutoMigrate()
	} else {
		G_db.CreateTable(&Class{})
	}
}

package main

import (
	"fmt"
	"time"

	"gitlab.com/indev-moph/fiber-api/helper"
)

type TestStruct struct {
	ID           uint       `json:"id" gorm:"primarykey" swaggerignore:"true"`
	Hospcode     string     `json:"hospcode" gorm:"column:hospcode;type:varchar(5)"`
	Username     string     `json:"username" gorm:"column:username;type:varchar(100)"`
	RDBMS        string     `json:"rdbms" gorm:"column:rdbms;type:varchar(50)"`
	OsServer     string     `json:"os_server" gorm:"column:os_server;type:varchar(50)"`
	Java         string     `json:"java" gorm:"column:java;type:varchar(10)"`
	ComputerName string     `json:"computer_name" gorm:"column:computer_name;type:varchar(500)"`
	OS           string     `json:"os" gorm:"column:os;type:varchar(50)"`
	JHCIS        string     `json:"jhcis" gorm:"column:jhcis;type:varchar(50)"`
	LocalIP      string     `json:"local_ip" gorm:"column:local_ip;type:varchar(50)"`
	PublicIP     string     `json:"public_ip" gorm:"column:public_ip;type:varchar(50)" swaggerignore:"true"`
	MacAddress   string     `json:"mac_address" gorm:"column:mac_address;type:varchar(30)"`
	ActiveAt     *time.Time `json:"active_at" gorm:"column:active_at"`
	UpdatedAt    time.Time  `json:"updated_at" swaggerignore:"true"`
	DeletedAt    *time.Time `json:"deleted_at" gorm:"index" swaggerignore:"true"`
}

func main() {
	s := helper.GetTagInStruct(TestStruct{}, "json")
	fmt.Println(s)

	s = helper.GetTagInStruct(&TestStruct{}, "json")
	fmt.Println(s)

	s = helper.GetTagInStruct(TestStruct{}, "gorm")
	fmt.Println(s)
}

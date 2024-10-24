package database

import (
	"shortLink/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	err error
)


func StartDB(){
	addr:= "host=172.20.0.2 user=root password=root dbname=root port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	DB,err = gorm.Open(postgres.Open(addr))
	if err != nil{
		panic(err.Error())
	}
	var Link models.Link
    DB.AutoMigrate(&Link)
}
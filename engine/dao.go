package engine

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var conn *gorm.DB

type Dao struct {
	Db *gorm.DB
}

func ConnectDb(dao *Dao) bool {
	if conn == nil {
		ok := initConnection()
		if !ok {
			return false
		}
	}
	dao.Db = conn
	return true
}

func initConnection() bool {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	if user == "" || password == "" {
		return false
	}
	if host == "" {
		host = "127.0.0.1"
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/vitroplus?charset=utf8mb4&parseTime=True&loc=Local", user, password, host)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		PrintError("Error connecting to db", err)
		return false
	}
	conn = db
	return true
}

func (dao *Dao) ExistsById(id uint, model interface{}) bool {
	var exists bool
	err := dao.Db.Model(model).Select("count(*)").Where("id = ?", id).Find(&exists).Error

	return exists && err == nil
}

func (dao *Dao) Create(model interface{}) error {
	result := dao.Db.Create(model)

	return result.Error
}

func (dao *Dao) Update(model interface{}, update interface{}) error {
	result := dao.Db.Model(model).Updates(update)

	return result.Error
}

func (dao *Dao) Get(id uint, model interface{}) error {
	result := dao.Db.First(model, id)

	return result.Error
}

func (dao *Dao) Delete(id uint, model interface{}) error {
	result := dao.Db.Delete(model, id)

	return result.Error
}

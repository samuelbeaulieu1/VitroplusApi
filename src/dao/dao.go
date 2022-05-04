package dao

import (
	"fmt"
	"os"

	"github.com/samuelbeaulieu1/gimlet"
	"github.com/samuelbeaulieu1/gimlet/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var conn *gorm.DB

type Dao struct {
	Db *gorm.DB
}

func connectDb(dao *Dao) bool {
	if conn == nil {
		ok := InitConnection()
		if !ok {
			return false
		}
	}
	dao.Db = conn
	return true
}

func InitConnection() bool {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	database := os.Getenv("DB")
	if user == "" || password == "" || database == "" {
		return false
	}
	if host == "" {
		host = "127.0.0.1"
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.PrintError("Error connecting to db %s", err)
		return false
	}
	conn = db
	return true
}

func (dao *Dao) ExistsByID(id string, model gimlet.Model) bool {
	var exists bool
	err := dao.Db.Model(model).Select("count(*) > 0").
		Where("id = ?", id).
		Find(&exists).Error

	return exists && err == nil
}

func (dao *Dao) Create(model gimlet.Model) error {
	result := dao.Db.Create(model)

	return result.Error
}

func (dao *Dao) Update(model gimlet.Model, update gimlet.Model) error {
	result := dao.Db.Model(model).Updates(update)

	return result.Error
}

func (dao *Dao) Get(id string, model gimlet.Model) error {
	result := dao.Db.Where("id = ?", id).First(model)

	return result.Error
}

func (dao *Dao) Delete(id string, model gimlet.Model) error {
	result := dao.Db.Where("id = ?", id).Delete(model)

	return result.Error
}

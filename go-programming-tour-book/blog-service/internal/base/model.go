package base

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"yazidchen.com/go-programming-tour-book/blog-service/global"
	"yazidchen.com/go-programming-tour-book/blog-service/pkg/setting"
)

type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedBy string `json:"modified_by"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

func NewDBEngine(d *setting.DatabaseSettingS) (*gorm.DB, error) {
	s := "%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Asia%%2FShanghai"
	db, err := gorm.Open(d.DBType, fmt.Sprintf(s,
		d.UserName,
		d.Password,
		d.Host,
		d.DBName,
		d.Charset,
		d.ParseTime,
	))
	if err != nil {
		return nil, err
	}

	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(d.MaxIdleConns)
	db.DB().SetMaxOpenConns(d.MaxOpenConns)

	return db, err
}

package datasource

import "github.com/jinzhu/gorm"

type IDb interface {
	DB() *gorm.DB
	Connect() error
}

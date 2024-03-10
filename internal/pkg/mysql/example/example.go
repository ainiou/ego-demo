package example

import (
	"context"
	"database/sql"
	"ego-demo/internal/pkg/mysql/db"
	"ego-demo/internal/pkg/mysql/option"
	"gorm.io/gorm"
)

type Example struct {
	*option.Option
}

func NewExampleInterface(opt *option.Option) *Example {
	return &Example{
		Option: opt,
	}
}

// BeginTransaction 开启事务
func (h *Example) BeginTransaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) error {
	return h.ExampleDB.Transaction(fc)
}

// GetRecord ...
func (h *Example) GetRecord(ctx context.Context) (record db.HelloWorld, err error) {
	if err = h.ExampleDB.WithContext(ctx).Model(new(db.HelloWorld)).Find(&record).Error; err != nil {
		return
	}
	return
}

// AddRecord ...
func (h *Example) AddRecord(ctx context.Context, record db.HelloWorld, _dbConn *gorm.DB) (err error) {
	var dbConn = h.ExampleDB
	if _dbConn != nil {
		dbConn = _dbConn
	}
	if err = dbConn.WithContext(ctx).Model(new(db.HelloWorld)).Create(&record).Error; err != nil {
		return
	}
	return
}

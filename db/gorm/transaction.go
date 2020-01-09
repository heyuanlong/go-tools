package gorm

import (
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Transaction struct {
	tx     *gorm.DB
	status int //0无事务，1事务中
}

func NewTransaction() *Transaction {
	return &Transaction{}
}
func (ts *Transaction) Begin(gormDB *gorm.DB) *gorm.DB {
	if ts.status == 0 {
		ts.tx = gormDB.Begin()
		ts.status = 1
		return ts.tx
	} else {
		return nil
	}
}

func (ts *Transaction) Commit() error {
	if ts.status == 1 {
		if err := ts.tx.Commit().Error; err != nil {
			return err
		}
		ts.status = 0
	} else {
		return errors.New("not transaction Commit")
	}
	return nil
}

func (ts *Transaction) Rollback() error {
	if ts.status == 1 {
		if err := ts.tx.Rollback().Error; err != nil {
			return err
		}
		ts.status = 0
	} else {
		return errors.New("not transaction Rollback")
	}
	return nil
}

func (ts *Transaction) Defer() error {
	if ts.status == 1 {
		if err := ts.tx.Commit().Error; err != nil {
			return err
		}
		ts.status = 0
	} else {
		return errors.New("not transaction Commit")
	}
	return nil
}

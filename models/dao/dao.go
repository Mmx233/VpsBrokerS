package dao

import "gorm.io/gorm"

func Begin() *gorm.DB {
	return db.Begin()
}

func defaultInsert(tx *gorm.DB, a interface{}) error {
	if tx == nil {
		tx = db
	}
	return tx.Create(a).Error
}

func defaultDelete(tx *gorm.DB, a interface{}) error {
	if tx == nil {
		tx = db
	}
	return tx.Where(a).Delete(a).Error
}

func defaultFind(a interface{}) error {
	return db.Where(a).Find(a).Error
}

func defaultExist(a interface{}) bool {
	var t struct {
		ID uint
	}
	db.Model(a).Where(a).Find(&t)
	return t.ID != 0
}

func defaultGet(t interface{}) error {
	return db.Find(t).Error
}

func defaultGetWhitQuery(a interface{}, t interface{}) error {
	return db.Where(a).Find(t).Error
}

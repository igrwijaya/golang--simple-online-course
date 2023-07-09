package db

import "gorm.io/gorm"

func Paginate(page int, size int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		currentPage := page

		if currentPage == 0 {
			currentPage = 1
		}

		if size == -1 {
			return db
		}

		offset := (currentPage - 1) * size

		return db.Offset(offset).Limit(size)
	}
}

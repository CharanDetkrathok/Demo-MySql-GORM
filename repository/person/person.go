package personrepo

import (
	"database/sql"
	"demo-mysql-gorm/domain"

	"gorm.io/gorm"
)

type personRepo struct {
	db      *sql.DB
	gorm_db *gorm.DB
}

func NewPersonRepo(db *sql.DB, gorm_db *gorm.DB) domain.PersonRepo {
	return &personRepo{
		db:      db,
		gorm_db: gorm_db,
	}
}

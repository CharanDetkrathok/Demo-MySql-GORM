package domain

import (
	"database/sql"
	"demo-mysql-gorm/model/entity"
)

type PersonRepo interface {
	GetPersonWithPersonID(personId int) (*entity.Person, error)
	GetPersonWithPersonID_GORM(personId int) (*entity.Person, error)

	InsertPerson(newPerson *entity.Person) (sql.Result, error)
	InsertPerson_GORM(newPerson *entity.Person) error

	UpdatePerson(newPerson *entity.Person) (sql.Result, error)
	UpdatePerson_GORM(newPerson *entity.Person) error

	DeletePerson(personID int) (sql.Result, error)
	DeletePerson_GORM(personID int) error
}

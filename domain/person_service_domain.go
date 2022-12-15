package domain

import (
	"demo-mysql-gorm/model/entity"
	"demo-mysql-gorm/model/request"
)

type PersonService interface {
	GetPersonWithPersonID(personId int) (*entity.Person, error)
	GetPersonWithPersonID_GORM(personId int) (*entity.Person, error)

	InsertPerson(newPerson *request.Person) error
	InsertPerson_GORM(newPerson *request.Person) error

	UpdatePerson(newPerson *request.Person) error
	UpdatePerson_GORM(newPerson *request.Person) error
}

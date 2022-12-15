package personrepo

import (
	"demo-mysql-gorm/model/entity"
	"fmt"
)

func (repo *personRepo) GetPersonWithPersonID(personId int) (*entity.Person, error) {
	var person entity.Person
	err := repo.db.QueryRow(getPersonWithPersonID, personId).Scan(
		&person.PersonID,
		&person.FirstName,
		&person.LastName,
		&person.Address,
		&person.City)
	if err != nil {
		fmt.Println("result error :: ", err)
		return &person, err
	}

	return &person, nil
}

func (repo *personRepo) GetPersonWithPersonID_GORM(personId int) (*entity.Person, error) {
	var person entity.Person
	err := repo.gorm_db.First(&person, personId).Error
	if err != nil {
		fmt.Println("GORM result error :: ", err)
		return &person, err
	}

	return &person, nil
}

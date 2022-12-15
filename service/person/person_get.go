package personservice

import (
	"demo-mysql-gorm/model/entity"
	"fmt"
)

func (service *personService) GetPersonWithPersonID(personId int) (*entity.Person, error) {

	person, err := service.PersonRepo.GetPersonWithPersonID(personId)
	if err != nil {
		return nil, fmt.Errorf("personID not found")
	}

	return person, nil

}

func (service *personService) GetPersonWithPersonID_GORM(personId int) (*entity.Person, error) {

	person, err := service.PersonRepo.GetPersonWithPersonID_GORM(personId)
	if err != nil {
		return nil, fmt.Errorf("gorm personID not found")
	}

	return person, nil

}

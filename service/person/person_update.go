package personservice

import (
	"demo-mysql-gorm/model/entity"
	"demo-mysql-gorm/model/request"
	"fmt"
)

func (service *personService) UpdatePerson(newPerson *request.Person) error {

	newPersonTemp := entity.Person{
		PersonID:  newPerson.PersonID,
		LastName:  newPerson.LastName,
		FirstName: newPerson.FirstName,
		Address:   newPerson.Address,
		City:      newPerson.City,
	}

	fmt.Println("newPersonTemp :: ", newPersonTemp)

	_, err := service.PersonRepo.GetPersonWithPersonID(newPerson.PersonID)
	if err != nil {
		fmt.Println("update SQL failed, not found person id. :: ", err)
		return fmt.Errorf("update SQL failed, not found person id")
	}

	result, err := service.PersonRepo.UpdatePerson(&newPersonTemp)
	if err != nil {
		fmt.Println("update SQL failed. :: ", err)
		return fmt.Errorf("update SQL failed")
	}

	fmt.Println("SUCCESS update by SQL :: ", result)
	return nil

}

func (service *personService) UpdatePerson_GORM(newPerson *request.Person) error {

	newPersonTemp := entity.Person{
		PersonID:  newPerson.PersonID,
		LastName:  newPerson.LastName,
		FirstName: newPerson.FirstName,
		Address:   newPerson.Address,
		City:      newPerson.City,
	}

	fmt.Println("newPersonTemp :: ", newPersonTemp)

	_, err := service.PersonRepo.GetPersonWithPersonID_GORM(newPerson.PersonID)
	if err != nil {
		fmt.Println("gorm update failed, not found person id. :: ", err)
		return fmt.Errorf("gorm update failed, not found person id")
	}

	err = service.PersonRepo.UpdatePerson_GORM(&newPersonTemp)
	if err != nil {
		fmt.Println("gorm update failed. :: ", err)
		return fmt.Errorf("gorm update failed")
	}

	fmt.Println("SUCCESS update by GORM")
	return nil

}

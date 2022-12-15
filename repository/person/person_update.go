package personrepo

import (
	"database/sql"
	"demo-mysql-gorm/model/entity"
)

func (repo *personRepo) UpdatePerson(newPerson *entity.Person) (sql.Result, error) {

	result, err := repo.db.Exec(updatePerson, newPerson.LastName, newPerson.FirstName, newPerson.Address, newPerson.City, newPerson.PersonID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (repo *personRepo) UpdatePerson_GORM(newPerson *entity.Person) error {

	result := repo.gorm_db.Model(&newPerson).Where("PersonID = ?", newPerson.PersonID).Updates(newPerson)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

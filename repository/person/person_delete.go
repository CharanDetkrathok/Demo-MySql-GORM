package personrepo

import (
	"database/sql"
	"demo-mysql-gorm/model/entity"
)

func (repo *personRepo) DeletePerson(personID int) (sql.Result, error) {

	result, err := repo.db.Exec(deletePerson, personID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (repo *personRepo) DeletePerson_GORM(personID int) error {

	result := repo.gorm_db.Where("PersonID = ?", personID).Delete(&entity.Person{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

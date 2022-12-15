package personrepo

import (
	"database/sql"
	"demo-mysql-gorm/model/entity"
)

func (repo *personRepo) InsertPerson(newPerson *entity.Person) (sql.Result, error) {

	result, err := repo.db.Exec(insertPerson, newPerson.PersonID, newPerson.LastName, newPerson.FirstName, newPerson.Address, newPerson.City)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (repo *personRepo) InsertPerson_GORM(newPerson *entity.Person) error {

	result := repo.gorm_db.Create(newPerson)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

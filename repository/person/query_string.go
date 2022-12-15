package personrepo

const (
	getPersonWithPersonID = `
		SELECT 
			PersonID,
			FirstName,
			LastName,
			Address,
			City
		FROM
		Persons 
		WHERE 
		PersonID = ?;`

	insertPerson = `
		INSERT INTO Persons (
			PersonID,
			LastName,
			FirstName,
			Address,
			City
		)
		VALUE(
			?,
			?,
			?,
			?,
			?
		);`

	updatePerson = `
		UPDATE 
			Persons 
		SET 
			LastName = ?,
			FirstName= ?,
			Address  = ?,
			City= ?
		WHERE
			PersonID = ?;`
)

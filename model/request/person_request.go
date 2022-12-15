package request

type (
	Person struct {
		PersonID  int    `json:"person_id"`
		LastName  string `json:"last_name"`
		FirstName string `json:"first_name"`
		Address   string `json:"address"`
		City      string `json:"city"`
	}
)

package personservice

import (
	"demo-mysql-gorm/domain"

	"github.com/go-redis/redis/v9"
)

type personService struct {
	Redis      *redis.Client
	PersonRepo domain.PersonRepo
}

func NewPersonService(Redis *redis.Client, PersonRepo domain.PersonRepo) domain.PersonService {
	return &personService{
		Redis:      Redis,
		PersonRepo: PersonRepo,
	}
}

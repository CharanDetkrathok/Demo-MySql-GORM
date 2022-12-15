package server

import (
	"database/sql"
	"demo-mysql-gorm/domain"

	"github.com/go-redis/redis/v9"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	personhandle "demo-mysql-gorm/handle/person"
	personrepo "demo-mysql-gorm/repository/person"
	personservice "demo-mysql-gorm/service/person"
)

type (
	server struct {
		PersonHandle domain.PersonHandle
	}

	Server interface {
		Router(app *fiber.App)
		personGroup(app *fiber.App)
	}
)

func NewServer(db *sql.DB, gorm_db *gorm.DB, redis *redis.Client) Server {

	newPersonRepo := personrepo.NewPersonRepo(db, gorm_db)
	newPersonService := personservice.NewPersonService(redis, newPersonRepo)
	newPersonHandle := personhandle.NewPersonHandle(newPersonService)

	return &server{newPersonHandle}
}

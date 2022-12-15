package personhandle

import (
	"demo-mysql-gorm/model/request"
	"demo-mysql-gorm/model/response"

	"github.com/gofiber/fiber/v2"
)

func (handle *personHandle) InsertPerson(c *fiber.Ctx) error {

	newPerson := new(request.Person)
	if err := c.BodyParser(newPerson); err != nil {
		res := response.Response{
			Code:         "-SQL001",
			Data:         "",
			HttpCode:     200,
			ErrorMessage: "Inser SQL Bad request.",
		}
		return c.JSON(res)
	}

	if err := handle.PersonService.InsertPerson(newPerson); err != nil {
		res := response.Response{
			Code:         "-SQL102",
			Data:         "",
			HttpCode:     200,
			ErrorMessage: "Insert SQL failed.",
		}
		return c.JSON(res)
	}

	res := response.Response{
		Code:         "SQL102",
		Data:         "",
		HttpCode:     200,
		ErrorMessage: "Insert SQL Successfully.",
	}

	return c.JSON(res)
}

func (handle *personHandle) InsertPerson_GORM(c *fiber.Ctx) error {

	newPerson := new(request.Person)
	if err := c.BodyParser(&newPerson); err != nil {
		res := response.Response{
			Code:         "-GORM001",
			Data:         "",
			HttpCode:     200,
			ErrorMessage: "Insert GORM Bad request.",
		}
		return c.JSON(res)
	}

	if err := handle.PersonService.InsertPerson_GORM(newPerson); err != nil {
		res := response.Response{
			Code:         "-GORM102",
			Data:         "",
			HttpCode:     200,
			ErrorMessage: "Insert GORM  Failed.",
		}
		return c.JSON(res)
	}

	res := response.Response{
		Code:         "GORM102",
		Data:         "",
		HttpCode:     200,
		ErrorMessage: "Insert GORM Successfully.",
	}

	return c.JSON(res)
}

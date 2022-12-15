package personhandle

import (
	"demo-mysql-gorm/model/request"
	"demo-mysql-gorm/model/response"

	"github.com/gofiber/fiber/v2"
)

func (handle *personHandle) UpdatePerson(c *fiber.Ctx) error {

	newPerson := new(request.Person)
	if err := c.BodyParser(newPerson); err != nil {
		res := response.Response{
			Code:         "-SQL001",
			Data:         "",
			HttpCode:     200,
			ErrorMessage: "Update SQL Bad request.",
		}
		return c.JSON(res)
	}

	if err := handle.PersonService.UpdatePerson(newPerson); err != nil {
		res := response.Response{
			Code:         "-SQL103",
			Data:         "",
			HttpCode:     200,
			ErrorMessage: "Update SQL failed.",
		}
		return c.JSON(res)
	}

	res := response.Response{
		Code:         "SQL103",
		Data:         "",
		HttpCode:     200,
		ErrorMessage: "Update SQL Successfully.",
	}

	return c.JSON(res)
}

func (handle *personHandle) UpdatePerson_GORM(c *fiber.Ctx) error {

	newPerson := new(request.Person)
	if err := c.BodyParser(&newPerson); err != nil {
		res := response.Response{
			Code:         "-GORM001",
			Data:         "",
			HttpCode:     200,
			ErrorMessage: "Update GORM Bad request.",
		}
		return c.JSON(res)
	}

	if err := handle.PersonService.UpdatePerson_GORM(newPerson); err != nil {
		res := response.Response{
			Code:         "-GORM103",
			Data:         "",
			HttpCode:     200,
			ErrorMessage: "Update GORM  Failed.",
		}
		return c.JSON(res)
	}

	res := response.Response{
		Code:         "GORM103",
		Data:         "",
		HttpCode:     200,
		ErrorMessage: "Update GORM Successfully.",
	}

	return c.JSON(res)
}

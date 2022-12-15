package personhandle

import (
	"demo-mysql-gorm/model/response"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (handle *personHandle) DeletePerson(c *fiber.Ctx) error {

	personId, _ := strconv.Atoi(c.Params("personId"))
	if err := handle.PersonService.DeletePerson(personId); err != nil {
		res := response.Response{
			Code:         "-SQL104",
			Data:         "",
			HttpCode:     200,
			ErrorMessage: "Delete SQL failed.",
		}
		return c.JSON(res)
	}

	res := response.Response{
		Code:         "SQL104",
		Data:         "",
		HttpCode:     200,
		ErrorMessage: "Delete SQL Successfully.",
	}

	return c.JSON(res)
}

func (handle *personHandle) DeletePerson_GORM(c *fiber.Ctx) error {

	personId, _ := strconv.Atoi(c.Params("personId"))
	if err := handle.PersonService.DeletePerson_GORM(personId); err != nil {
		res := response.Response{
			Code:         "-GORM104",
			Data:         "",
			HttpCode:     200,
			ErrorMessage: "Delete GORM  Failed.",
		}
		return c.JSON(res)
	}

	res := response.Response{
		Code:         "GORM104",
		Data:         "",
		HttpCode:     200,
		ErrorMessage: "Delete GORM Successfully.",
	}

	return c.JSON(res)
}

package personhandle

import (
	"demo-mysql-gorm/model/response"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (handle *personHandle) GetPersonWithPersonID(c *fiber.Ctx) error {

	personId, _ := strconv.Atoi(c.Params("personId"))
	result, err := handle.PersonService.GetPersonWithPersonID(personId)
	if err != nil {
		res := response.Response{
			Code:         "-SQL101",
			Data:         "",
			HttpCode:     200,
			ErrorMessage: "Get SQL Failed.",
		}
		return c.JSON(res)
	}

	resp := response.Response{
		Code:         "SQL101",
		Data:         result,
		HttpCode:     200,
		ErrorMessage: "Get SQL Successfully.",
	}

	return c.JSON(resp)
}

func (handle *personHandle) GetPersonWithPersonID_GORM(c *fiber.Ctx) error {

	personId, _ := strconv.Atoi(c.Params("personId"))
	result, err := handle.PersonService.GetPersonWithPersonID_GORM(personId)
	if err != nil {
		res := response.Response{
			Code:         "-GROM101",
			Data:         "",
			HttpCode:     200,
			ErrorMessage: "Get GORM Failed.",
		}
		return c.JSON(res)
	}

	resp := response.Response{
		Code:         "GORM101",
		Data:         result,
		HttpCode:     200,
		ErrorMessage: "Get GORM Successfully.",
	}

	return c.JSON(resp)
}

package auth

import (
	"github.com/gofiber/fiber/v2"
	"mp_update_server_go/core/database"
	"mp_update_server_go/core/models/requests"
)

func Login(c *fiber.Ctx) error {
	var isExist int
	cred := &requests.LoginRequest{}
	if err := c.BodyParser(cred); err != nil {
		return err
	}

	db := database.InitializeDB()

	row := db.DbInstance.QueryRow("select count(id) from profile where username = $1 and password = $2 limit 1", cred.Username, cred.Password)
	row.Scan(&isExist)

	return nil
}

func UserInfo(c *fiber.Ctx) error {
	return nil
}

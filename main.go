package main

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/mailru/easyjson"
	"log"
	"mp_update_server_go/core/api/auth"
	"mp_update_server_go/core/api/updates"
	"mp_update_server_go/core/database"
	"mp_update_server_go/core/models/dao"
)

type VersionResponse struct {
	LatestVersion     string `json:"latestVersion"`
	LatestVersionCode int    `json:"latestVersionCode"`
	Url               string `json:"url"`
}

func getVersionOfApp(c *fiber.Ctx) error {
	appName := c.Params("appName")

	db := database.InitializeDB()

	version := &[]dao.Version{}

	if err := db.DbInstance.Select(version, "select * from version where app_id = $1 order by version_code desc limit 1", appName); err != nil {
		return err
	}

	if len(*version) == 0 {
		c.Status(404).JSON(map[string]interface{}{})
	}

	if err := c.JSON((*version)[0]); err != nil {
		return err
	}

	return nil
}

func main() {
	app := fiber.New(fiber.Config{JSONEncoder: func(v interface{}) ([]byte, error) {
		switch x := v.(type) {
		case easyjson.Marshaler:
			return easyjson.Marshal(x)
		default:
			return json.Marshal(x)
		}
	},

		JSONDecoder: func(data []byte, v interface{}) error {
			switch x := v.(type) {
			case easyjson.Unmarshaler:
				return easyjson.Unmarshal(data, x)
			default:
				return json.Unmarshal(data, x)
			}
		},
	})
	apiGroup := app.Group("/api")
	updates.InitRoutes(apiGroup)
	auth.InitRoutes(apiGroup)

	app.Get("/updates/:appName/update-changelog.json", getVersionOfApp)
	log.Fatal(app.Listen(":8001"))
}

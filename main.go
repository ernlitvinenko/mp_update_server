package main

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/mailru/easyjson"
	"log"
	"mp_update_server_go/core/api/auth"
	"mp_update_server_go/core/api/updates"
	"mp_update_server_go/core/database"
	"mp_update_server_go/core/models/dao"
	"mp_update_server_go/core/storage/s3"
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
	s3.New()

	app := fiber.New(fiber.Config{
		BodyLimit: 512 * 1024 * 1024,
		JSONEncoder: func(v interface{}) ([]byte, error) {
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

	app.Use(cors.New(cors.Config{AllowOrigins: "*"}))

	apiGroup := app.Group("/api")
	updates.InitRoutes(apiGroup)
	auth.InitRoutes(apiGroup)

	app.Get("/updates/:appName/update-changelog.json", getVersionOfApp)
	log.Fatal(app.Listen(":8001"))
}

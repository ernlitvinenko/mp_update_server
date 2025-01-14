package updates

import (
	"github.com/gofiber/fiber/v2"
	"mp_update_server_go/core/database"
	"mp_update_server_go/core/models/dao"
	"mp_update_server_go/core/models/requests"
)

// CreateApplication add app into database
func CreateApplication(c *fiber.Ctx) error {

	data := &requests.CreateApplicationRequest{}

	if err := c.BodyParser(data); err != nil {
		return err
	}

	db := database.InitializeDB()

	if _, err := db.DbInstance.Exec("call create_app($1, $2, $3)", data.Id, data.AppName, data.Link); err != nil {
		return err
	}

	return nil
}

func AddVersion(c *fiber.Ctx) error {
	data := &requests.AddVersionRequest{}
	if err := c.BodyParser(data); err != nil {
		return err
	}
	db := database.InitializeDB()

	if _, err := db.DbInstance.Exec("call add_version($1, $2, $3)", data.AppId, data.Id, data.Link); err != nil {
		return err
	}

	return nil
}

func ListApplications(c *fiber.Ctx) error {
	db := database.InitializeDB()
	data := &[]dao.ListApplicationDao{}

	if err := db.DbInstance.Select(data, "select a.id as id, a.app_name as name, v.id as version_id, v.version_code as version_code, v.link as version_link from application a left outer join version v on a.id = v.app_id;"); err != nil {
		return err
	}

	applications := map[string]dao.Application{}

	for _, val := range *data {
		if entry, ok := applications[val.Id]; ok {
			entry.Versions = append(entry.Versions, dao.Version{
				Id:          val.VersionId,
				VersionCode: val.VersionCode,
				Description: nil,
				Link:        val.Link,
			})

			applications[val.Id] = entry
			continue
		}
		applications[val.Id] = dao.Application{
			Id:      val.Id,
			AppName: val.AppName,
			Versions: []dao.Version{
				{Id: val.VersionId, VersionCode: val.VersionCode, Description: nil, Link: val.Link},
			},
		}
	}
	c.JSON(applications)

	return nil
}

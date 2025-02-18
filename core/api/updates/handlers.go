package updates

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"mime/multipart"
	"mp_update_server_go/core/database"
	"mp_update_server_go/core/models/dao"
	"mp_update_server_go/core/models/requests"
	"mp_update_server_go/core/models/responses"
	"mp_update_server_go/core/storage/s3"
	"net/http"
)

type UploadAppResponse struct {
	Link string `json:"link"`
}

func uploadFileToMinio(file *multipart.FileHeader, filename string) (*minio.UploadInfo, error) {
	ctx := context.Background()
	storage := s3.New()
	reader, err := file.Open()
	if err != nil {
		return nil, err
	}

	info, err := storage.Client.PutObject(ctx, "mp-update", filename, reader, file.Size, minio.PutObjectOptions{})

	return &info, err

}

func UploadApp(c *fiber.Ctx) error {
	file, err := c.FormFile("version")
	if err != nil || file == nil {
		return err
	}
	fileName := uuid.NewString() + file.Filename
	storage := s3.New()

	_, err = uploadFileToMinio(file, fileName)
	if err != nil {
		return err
	}

	err = c.JSON(UploadAppResponse{
		Link: "http://" + storage.Client.EndpointURL().Hostname() + ":" + storage.Client.EndpointURL().Port() + "/mp-update/" + fileName,
	})

	if err != nil {
		return err
	}

	return nil
}

// CreateApplication add app into database
func CreateApplication(c *fiber.Ctx) error {

	data := &requests.CreateApplicationRequest{}

	if err := c.BodyParser(data); err != nil {
		return err
	}

	db := database.InitializeDB()

	if _, err := db.DbInstance.Exec("call create_app($1, $2, $3, $4, $5)", data.Id, data.AppName, data.Link, data.Description, data.Version); err != nil {
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

	if _, err := db.DbInstance.Exec("call add_version($1, $2, $3, $4)", data.AppId, data.Id, data.Link, data.Description); err != nil {
		return err
	}

	return nil
}

func ListApplications(c *fiber.Ctx) error {
	db := database.InitializeDB()
	data := &[]dao.ListApplicationDao{}

	if err := db.DbInstance.Select(data, "select a.id as id, a.app_name as name, a.description as description, v.id as version_id, v.version_code as version_code, v.link as version_link, v.description as version_description from application a left outer join version v on a.id = v.app_id;"); err != nil {
		return err
	}

	applications := map[string]dao.Application{}

	for _, val := range *data {
		if entry, ok := applications[val.Id]; ok {
			entry.Versions = append(entry.Versions, dao.Version{
				Id:          val.VersionId,
				VersionCode: val.VersionCode,
				Description: val.VersionDescription,
				Link:        val.Link,
			})

			applications[val.Id] = entry
			continue
		}
		applications[val.Id] = dao.Application{
			Id:          val.Id,
			AppName:     val.AppName,
			Description: val.Description,
			Versions: []dao.Version{
				{Id: val.VersionId, VersionCode: val.VersionCode, Description: val.VersionDescription, Link: val.Link},
			},
		}
	}

	c.JSON(applications)

	return nil
}

func DeleteVersion(c *fiber.Ctx) error {
	db := database.InitializeDB()
	app := c.Params("appName")
	version := c.Params("version")

	querySearch := "select id, app_id from version where id = $1 and app_id = $2"

	rows, err := db.DbInstance.Query(querySearch, version, app)

	if err != nil {
		return err
	}

	if !(rows.Next()) {
		c.Status(http.StatusNotFound)
		c.JSON(responses.ErrorResponse{
			Status: http.StatusNotFound,
			Error:  "Версия не найдена",
		})
		return nil
	}

	query := "delete from version where id = $1 and app_id = $2"

	_, err = db.DbInstance.Exec(query, version, app)

	if err != nil {
		return err
	}

	c.JSON(map[string]interface{}{"application": app, "version": version})
	c.Status(http.StatusNoContent)

	return nil
}

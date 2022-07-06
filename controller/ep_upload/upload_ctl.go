package ep_upload

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/phuslu/log"
	"gitlab.com/indev-moph/fiber-api/route/regisroute"
	"gitlab.com/indev-moph/fiber-api/route/restful"
)

const (
	DIR_STORAGE = "./storage/upload"
)

type EndpointUploadFile struct{}

type ListFileInfo struct {
	Name     string    `json:"name"`     // ชื่อไฟล์
	Datetime time.Time `json:"datetime"` // วันที่ Upload
	Size     int64     `json:"size"`     // ขนาดข้อมูล Byte(B)
	SizeMB   string    `json:"size_mb"`  // ขนาดข้อมูล MagaByte(MB)
} // @name FileInfo

func Init() {
	_ = os.MkdirAll(DIR_STORAGE, os.ModePerm)

	regisroute.AddRoute(endpointUpload)

}

func endpointUpload(r fiber.Router) {
	endpoint := EndpointUploadFile{}
	rg := restful.NewRestful(r, "/uploads", endpoint, MiddlewareVerifyACL)
	rg.Name("uploads")
}

func MiddlewareVerifyACL(c *fiber.Ctx) error {
	username, ok := c.Locals("username").(string)
	if !ok || username != "admin" {
		return &fiber.Error{
			Code:    fiber.StatusForbidden,
			Message: "Permission Denind.",
		}
	}
	return c.Next()
}

// DownloadFile godoc
// @Summary  ListFiles
// @Description
// @Tags      Upload
// @Produce   json
// @Success   200      {object}  []ListFileInfo  list  file  upload
// @Failure   default  {string}  string
// @security  BasicAuth
// @Router    /uploads [get]
func (EndpointUploadFile) Get(c *fiber.Ctx) error {
	files, err := ioutil.ReadDir(DIR_STORAGE)
	if err != nil {
		return err
	}
	result := []ListFileInfo{}
	for _, f := range files {
		result = append(result, ListFileInfo{
			Name:     f.Name(),
			Datetime: f.ModTime(),
			Size:     f.Size(),
			SizeMB:   fmt.Sprintf("%.2f", (float64(f.Size()) / (1024.0 * 1024.0))),
		})
	}
	return c.JSON(result)
}

// DownloadFile godoc
// @Summary  DownloadFile
// @Description
// @Tags      Upload
// @Produce   octet-stream
// @Param     name     path      string  true  "Filename for download."
// @Success   200      {string}  string
// @Failure   default  {string}  string
// @security  BasicAuth
// @Router    /uploads/{name} [get]
func (EndpointUploadFile) GetByID(c *fiber.Ctx, name string) error {
	filename := path.Join(DIR_STORAGE, name)
	if _, err := os.Lstat(filename); err != nil {
		log.Debug().Str("filename", filename).Msg("File Not Found.")
		return c.Status(404).SendString("File Not Found.")
	}
	return c.SendFile(filename, false)
}

// UploadFile godoc
// @Summary  UploadFile
// @Description
// @Tags      Upload
// @Produce   octet-stream
// @Param     n        formData  string  false  "Set filename (Default: fileupload name)"
// @Param     f        formData  file    true   "FileUpload."
// @Success   200      {string}  string
// @Failure   default  {string}  string
// @security  BasicAuth
// @Router    /uploads [post]
func (EndpointUploadFile) Add(c *fiber.Ctx) error {
	file, err := c.FormFile("f")
	if err != nil {
		return err
	}
	filename := c.FormValue("n", file.Filename)
	log.Debug().Msgf("Upload File: %s", filename)
	return c.SaveFile(file, path.Join(DIR_STORAGE, filename))
}

// DeleteFile godoc
// @Summary  DeleteFile
// @Description
// @Tags      Upload
// @Produce   plain
// @Param     name     path      string  true  "Filename for delete."
// @Success   200      {string}  string
// @Failure   default  {string}  string
// @security  BasicAuth
// @Router    /uploads/{name} [delete]
func (EndpointUploadFile) Delete(c *fiber.Ctx, name string) error {
	filename := path.Join(DIR_STORAGE, name)
	if _, err := os.Stat(filename); err != nil {
		return c.Status(404).SendString("File Not Found.")
	}
	return os.Remove(filename)
}

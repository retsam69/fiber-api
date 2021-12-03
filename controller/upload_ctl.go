package controller

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"os"
// 	"path"
// 	"time"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/phuslu/log"
// )

// const (
// 	DIR_STORAGE = "./storage/upload"
// )

// type ListFileInfo struct {
// 	Name     string    `json:"name"`     // ชื่อไฟล์
// 	Datetime time.Time `json:"datetime"` // วันที่ Upload
// 	Size     int64     `json:"size"`     // ขนาดข้อมูล Byte(B)
// 	SizeMB   string    `json:"size_mb"`  // ขนาดข้อมูล MagaByte(MB)
// } // @name FileInfo

// func init() {
// 	_ = os.MkdirAll(DIR_STORAGE, os.ModePerm)

// 	RegisRoutes = append(RegisRoutes, func(r fiber.Router) {
// 		r.All("/ul/:name?/*", EndpointUpload)
// 	})
// }

// func EndpointUpload(c *fiber.Ctx) error {
// 	username := c.Locals("username").(string)
// 	if username != "admin" {
// 		return c.SendStatus(403)
// 	}
// 	name := fmt.Sprintf("%s /%s", c.Method(), c.Params("name", ""))
// 	switch name {
// 	case "GET /list":
// 		return ListFile(c)
// 	case "GET /dl":
// 		return DownloadFile(c)
// 	case "POST /":
// 		return UploadFile(c)
// 	case "DELETE /":
// 		return DeleteFile(c)
// 	}
// 	return c.SendStatus(404)
// }

// // ListFiles godoc
// // @Summary  ListFiles
// // @Description
// // @Tags      Upload
// // @Produce   json
// // @Success   200      {object}  []ListFileInfo  status
// // @Failure   default  {string}  string
// // @security  BasicAuth
// // @Router    /ul/list [get]
// func ListFile(c *fiber.Ctx) error {
// 	files, err := ioutil.ReadDir(DIR_STORAGE)
// 	if err != nil {
// 		return err
// 	}
// 	result := []ListFileInfo{}
// 	for _, f := range files {
// 		result = append(result, ListFileInfo{
// 			Name:     f.Name(),
// 			Datetime: f.ModTime(),
// 			Size:     f.Size(),
// 			SizeMB:   fmt.Sprintf("%.2f", (float64(f.Size()) / (1024.0 * 1024.0))),
// 		})
// 	}
// 	return c.JSON(result)
// }

// // DownloadFile godoc
// // @Summary  DownloadFile
// // @Description
// // @Tags      Upload
// // @Produce   octet-stream
// // @Param     f        query     string  true  "Filename for download."
// // @Success   200      {string}  string
// // @Failure   default  {string}  string
// // @security  BasicAuth
// // @Router    /ul/{dl} [get]
// func DownloadFile(c *fiber.Ctx) error {
// 	fdl := c.Query("f", ".emptyfile")
// 	filename := path.Join(DIR_STORAGE, fdl)
// 	if _, err := os.Lstat(filename); err != nil {
// 		log.Debug().Str("filename", filename).Msg("File Not Found.")
// 		return c.Status(404).SendString("File Not Found.")
// 	}
// 	return c.SendFile(filename, false)
// }

// // UploadFile godoc
// // @Summary  UploadFile
// // @Description
// // @Tags      Upload
// // @Produce   octet-stream
// // @Param     n        formData  string  false  "Set filename (Default: fileupload name)"
// // @Param     f        formData  file    true   "FileUpload."
// // @Success   200      {string}  string
// // @Failure   default  {string}  string
// // @security  BasicAuth
// // @Router    /ul/dl [post]
// func UploadFile(c *fiber.Ctx) error {
// 	file, err := c.FormFile("f")
// 	if err != nil {
// 		return err
// 	}
// 	filename := c.FormValue("n", file.Filename)
// 	log.Debug().Msgf("Upload File: %s", filename)
// 	return c.SaveFile(file, path.Join(DIR_STORAGE, filename))
// }

// // DeleteFile godoc
// // @Summary  DeleteFile
// // @Description
// // @Tags      Upload
// // @Produce   plain
// // @Param     f        query     string  true  "Filename for delete."
// // @Success   200      {string}  string
// // @Failure   default  {string}  string
// // @security  BasicAuth
// // @Router    /ul/dl [delete]
// func DeleteFile(c *fiber.Ctx) error {
// 	fdl := c.Query("f", "")
// 	filename := path.Join(DIR_STORAGE, fdl)
// 	if _, err := os.Stat(filename); err != nil {
// 		return c.Status(404).SendString("File Not Found.")
// 	}
// 	return os.Remove(filename)
// }

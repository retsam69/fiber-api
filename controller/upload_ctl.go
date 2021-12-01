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
// 	Name     string    `json:"name"`
// 	Datetime time.Time `json:"datetime"`
// 	Size     int64     `json:"size"`
// 	SizeMB   string    `json:"size_mb"`
// }

// func init() {
// 	_ = os.MkdirAll(DIR_STORAGE, os.ModePerm)

// 	RegisRoutes = append(RegisRoutes, func(r fiber.Router) {
// 		r.All("/admin/:name?/*", EndpointUpload)
// 	})
// }

// func EndpointUpload(c *fiber.Ctx) error {

// 	username := c.Locals("username").(string)
// 	if username != "admin" {
// 		return c.SendStatus(403)
// 	}

// 	name := c.Params("name", "")
// 	switch c.Method() {
// 	case "GET":
// 		if name == "list" {
// 			return ListFile(c)
// 		} else if name == "dl" {
// 			return DownloadFile(c)
// 		}
// 	case "POST":
// 		return UploadFile(c)
// 	case "DELETE":
// 		return DeleteFile(c)
// 	}
// 	return c.SendStatus(404)
// }

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

// func DownloadFile(c *fiber.Ctx) error {
// 	fdl := c.Query("f", ".emptyfile")
// 	filename := path.Join(DIR_STORAGE, fdl)
// 	if _, err := os.Lstat(filename); err != nil {
// 		log.Debug().Str("filename", filename).Msg("File Not Found.")
// 		return c.Status(404).SendString("File Not Found.")
// 	}
// 	return c.SendFile(filename, false)
// }

// func UploadFile(c *fiber.Ctx) error {
// 	file, err := c.FormFile("f")
// 	if err != nil {
// 		return err
// 	}
// 	filename := c.FormValue("n", file.Filename)
// 	log.Debug().Msgf("Upload File: %s", filename)
// 	return c.SaveFile(file, path.Join(DIR_STORAGE, filename))
// }

// func DeleteFile(c *fiber.Ctx) error {
// 	fdl := c.Query("f", "")
// 	filename := path.Join(DIR_STORAGE, fdl)
// 	if _, err := os.Stat(filename); err != nil {
// 		return c.Status(404).SendString("File Not Found.")
// 	}
// 	return os.Remove(filename)
// }

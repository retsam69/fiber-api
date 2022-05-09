package route

// import (
// 	"fmt"

// 	rapidoc "github.com/attapon-th/gofiber-rapidoc"
// 	"github.com/gofiber/fiber/v2"
// 	"github.com/spf13/viper"
// )

// func EndpointSwagger(r fiber.Router, prefix string) {
// 	UrlSwaggerFile := fmt.Sprintf("%s/docs/openapi-%s.json", prefix, viper.GetString("version"))
// 	r.Get(prefix+"/*", rapidoc.New(rapidoc.Config{
// 		Title:       "Service API",
// 		HeaderText:  "Service API",
// 		RenderStyle: rapidoc.RenderStyle_Read,
// 		SchemaStyle: rapidoc.SchemaStyle_Table,
// 		SpecURL:     UrlSwaggerFile,
// 		LogoURL:     `https://cdn-icons-png.flaticon.com/512/2165/2165004.png`,
// 	}))

// }

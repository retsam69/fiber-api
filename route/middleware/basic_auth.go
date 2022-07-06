package middleware

// import (
// "github.com/gofiber/fiber/v2/middleware/basicauth"
// )
// * ►─═ Middleware Http Basic Auth ═─►
// func BasicAuth() fiber.Handler {
// 	users := viper.GetString("auth.basic")
// 	var u = make(map[string]string)
// 	for _, v := range strings.Split(users, ",") {
// 		// log.Debug().Msg(v)
// 		t := strings.SplitN(v, ":", 2)
// 		if len(t) == 2 {
// 			u[strings.TrimSpace(t[0])] = strings.TrimSpace(t[1])
// 			log.Debug().Str("Username", strings.TrimSpace(t[0])).Msgf("BasicAuth: User Registered")
// 		}
// 	}
// 	return basicauth.New(basicauth.Config{
// 		Users: u,
// 	})

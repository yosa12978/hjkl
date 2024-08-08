package main

import (
	"github.com/yosa12978/hjkl/app"
	_ "github.com/yosa12978/hjkl/docs"
)

//	@title			HJKL API
//	@version		0.1
//	@description	Just an api.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	MIT

//	@host		localhost:5000
//	@BasePath	/api

//	@securityDefinitions.basic	BasicAuth

//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	if err := app.Run(); err != nil {
		panic(err)
	}
}

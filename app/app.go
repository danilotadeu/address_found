package app

import (
	"log"
	"os"

	"github.com/danilotadeu/r-customer-code-information/app/codeinformation"
)

//Container ...
type Container struct {
	CodeInformation codeinformation.App
}

//Register app container
func Register() *Container {
	container := &Container{
		CodeInformation: codeinformation.NewApp(os.Getenv("URL_PROVIDER"), os.Getenv("PORT_PROVIDER")),
	}
	log.Println("Registered -> App")
	return container
}

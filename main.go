package main

import (
	"fmt"
	"github.com/sandeepmendiratta/restapi-go/service"
)

var appName = "respapigo"

func main() {

	fmt.Printf("Starting %v\n", appName)
	service.StartWebServer("6767")
}



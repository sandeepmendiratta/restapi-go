package main

import (
	"github.com/sandeepmendiratta/restapi-go/service"
)

//var appName = "respapigo"

func main() {
	//fmt.Printf("Starting %v\n", appName)
	service.StartWebServer("6767")
}



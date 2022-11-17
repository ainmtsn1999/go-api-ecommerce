package main

import (
	"fmt"

	"github.com/ainmtsn1999/go-api-ecommerce/config"
	"github.com/ainmtsn1999/go-api-ecommerce/db"
	"github.com/ainmtsn1999/go-api-ecommerce/routers"
)

func main() {

	//koneksi database
	db.ConnectDB()

	//router
	e := routers.Init()
	port := fmt.Sprintf(":%s", config.GetEnvVariable("APP_PORT"))
	e.Logger.Fatal(e.Start(port))
}

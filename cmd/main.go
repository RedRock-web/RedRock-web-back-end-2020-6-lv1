package main

import (
	"RedRock-web-back-end-2020-6-lv1/app"
	"RedRock-web-back-end-2020-6-lv1/database"
	"RedRock-web-back-end-2020-6-lv1/router"
)

func main() {
	database.Start()
	app.Start()
	router.Start()
}

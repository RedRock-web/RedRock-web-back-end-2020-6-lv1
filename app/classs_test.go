package app

import (
	"RedRock-web-back-end-2020-6-lv1/database"
	"testing"
)

func TestGetAllClassInfo(t *testing.T) {
	database.Start()
	defer GetAllClassInfo(2019211555)
	recover()
}
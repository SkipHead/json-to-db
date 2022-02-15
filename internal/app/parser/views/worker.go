package views

import (
	"fmt"
	"json-to-postgres/internal/app/parser/controllers"
	"json-to-postgres/internal/app/parser/models"
)

func Work(path string) models.JsonStruct {
	j := controllers.Parse(controllers.OpenFileJson(path))
	fmt.Println(j)
	return j
}

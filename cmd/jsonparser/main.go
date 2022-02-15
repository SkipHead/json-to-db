package main

import (
	"fmt"
	"json-to-postgres/internal/app/parser/controllers"
	"json-to-postgres/internal/app/parser/views"
	"os"
)

func main() {

	file := os.Args[1:]
	fmt.Println(os.Args[1:])
	controllers.WriteDataDB(views.Work(file[0]))
}

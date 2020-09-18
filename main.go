package main

import (
	gin "cms/src/routes"
	"os"
)

func main() {
	var r gin.Routes
	router := r.StartGin()

	router.Run(os.Getenv("PORT"))
}

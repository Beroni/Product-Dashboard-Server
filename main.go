package main

import (
	gin "cms/src/routes"
)

func main() {
	var r gin.Routes
	router := r.StartGin()
	router.Run(":3333")
}

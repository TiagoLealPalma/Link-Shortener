package main

import (
	"Link-Shortener/data"
	"Link-Shortener/routes"
	"fmt"
)

func main() {
	data.InitDB()

	link, err := data.GetOriginalLink("abc123")
	if err != nil {
		panic(err)
	}

	fmt.Println("Link original:", link)

	r := routes.SetupRouter()

	r.StaticFile("/", "./index.html")

	r.Run(":8080")

}

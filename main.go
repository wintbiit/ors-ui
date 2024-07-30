package main

import "github.com/wintbiit/ors-ui/router"

func main() {
	r := router.NewRouter()

	r.Run(":8080")
}

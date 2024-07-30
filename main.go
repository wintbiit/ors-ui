package main

import (
	"github.com/wintbiit/ors-ui/internal"
	"github.com/wintbiit/ors-ui/router"
)

func main() {
	r := router.NewRouter()

	r.Run(internal.Config.Addr)
}

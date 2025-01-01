package main

import (
	"fmt"

	"github.com/rudransh-shrivastava/self-space/app"
	"github.com/rudransh-shrivastava/self-space/db"
)

func main() {
	db, err := db.NewDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	apiServer := app.NewApiServer(db)
	apiServer.Start()
}

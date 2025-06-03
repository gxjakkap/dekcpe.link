package main

import (
	"fmt"

	"github.com/gxjakkap/dekcpe.link/db"
	"github.com/gxjakkap/dekcpe.link/handler"
	"github.com/gxjakkap/dekcpe.link/router"
	"github.com/gxjakkap/dekcpe.link/store"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Warning: No .env file found")
	}

	r := router.New()
	d := db.New()

	cs := store.NewClicksStore(d)
	ls := store.NewLinkStore(d)

	h := handler.NewHandler(ls, cs)
	h.Register(r)
	err = r.Listen(":3000")
	if err != nil {
		fmt.Printf("%v", err)
	}
}

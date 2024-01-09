package main

import "github.com/emejotaw/voting-service/config"

func main() {

	db := config.NewDatabase()

	_ = db
}

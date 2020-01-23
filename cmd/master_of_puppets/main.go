package main

import "github.com/xcorter/master_of_puppets/internal/app/master_of_puppets"

func main() {
	app := master_of_puppets.NewApp()
	app.Start()
}

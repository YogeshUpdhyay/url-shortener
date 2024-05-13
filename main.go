package main

import (
	app "github.com/YogeshUpdhyay/url-shortener/internal"
)

func main() {
	app := app.GetApp()

	app.Run()
}

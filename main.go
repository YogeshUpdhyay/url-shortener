package main

import (
	app "github.com/YogeshUpdhyay/url-shortner/internal"
)

func main() {
	app := app.GetApp()

	app.Run()
}

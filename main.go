package main

import (
	app "url-shortner/internal"
)

func main() {
	app := app.GetApp()

	app.Run()
}

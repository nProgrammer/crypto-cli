package main

import (
	"crypto-cli/controllers"
)

func main() {
	a, crypto := controllers.PrepareApp()
	controllers.App(a, crypto)
}

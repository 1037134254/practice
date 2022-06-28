package main

import "example.com/m/v2/router"

func main() {
	engine := router.Router()
	err := engine.Run()
	if err != nil {
		return
	}
}

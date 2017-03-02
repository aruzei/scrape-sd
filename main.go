package main

import (
	"log"

	"github.com/sclevine/agouti"
)

func main() {
	driver := agouti.ChromeDriver()
	if err := driver.Start(); err != nil {
		log.Fatalf("Failed to start driver:%v", err)
	}
	defer driver.Stop()

	page, err := driver.NewPage(agouti.Browser("chrome"))
	if err != nil {
		log.Fatalf("Failed to open page:%v", err)
	}

	if err := page.Navigate("http://play.pokemonshowdown.com/battles"); err != nil {
		log.Fatalf("Failed to navigate:%v", err)
	}
	page.Screenshot("./.debug/debug.jpg")
}

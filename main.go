package main

import "scrape-sd.local/browser"

func main() {
	browser := browser.CreateChrome()
	defer browser.Stop()

	browser.NavigatePage("http://play.pokemonshowdown.com/battles")

	browser.Screenshot("./.debug/debug.jpg")
}

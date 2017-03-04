package browser

import (
	"log"

	"github.com/sclevine/agouti"
)

// CreateChrome create google chrome's browser
func CreateChrome() Browser {
	d, p := createChrome()
	browser := newBrowser(d, p)
	return browser
}

func createChrome() (*agouti.WebDriver, *agouti.Page) {
	driver := agouti.ChromeDriver()
	if err := driver.Start(); err != nil {
		log.Fatalf("Failed to start driver:%v", err)
	}
	page, err := driver.NewPage(agouti.Browser("chrome"))
	if err != nil {
		log.Fatalf("Failed to open page:%v", err)
	}
	page.Size(512, 512)
	return driver, page
}

package browser

import (
	"log"

	"github.com/sclevine/agouti"
)

// Browser contains webDriver and page
type Browser struct {
	*agouti.WebDriver
	*agouti.Page
}

// NavigatePage navigate to url.
// This wraps agouti.Page.Navigate(url string) error
func (browser *Browser) NavigatePage(url string) {
	if err := browser.Navigate(url); err != nil {
		log.Fatalf("Failed to navigate:%v", err)
	}
}

func newBrowser(w *agouti.WebDriver, p *agouti.Page) Browser {
	return Browser{w, p}
}

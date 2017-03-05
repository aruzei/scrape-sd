package browser

import (
	"log"
	"time"

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

//ExecuteWithWait execute fucntion affter three seconds.
func (browser *Browser) ExecuteWithWait(f func() error) error {
	time.Sleep(3 * time.Second)
	return f()
}

//WaitSeconds wait seconds
func (browser *Browser) WaitSeconds(seconds int) error {
	for index := 0; index < seconds; index++ {
		time.Sleep(time.Duration(1 * time.Second))
	}
	return nil
}

// ClickElement clicks a specified element with three second wait.
func (browser *Browser) ClickElement(element *agouti.Selection) error {
	return browser.ExecuteWithWait(element.Click)
}

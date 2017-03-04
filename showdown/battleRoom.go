package showdown

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"scrape-sd.local/browser"
)

type BattleRoom struct {
	Browser *browser.Browser
}

func NewBattleRoom() *BattleRoom {
	browser := browser.CreateChrome()
	doNothing := func() error { return nil }
	browser.ExecuteWithWait(doNothing)

	browser.NavigatePage("http://play.pokemonshowdown.com/battles")

	room := BattleRoom{
		Browser: &browser}

	return &room
}

// Scrape scrapes html in http://play.pokemonshowdown.com/battles
func (room *BattleRoom) Scrape() []BattleLink {

	browser := room.Browser

	browser.ClickElement(browser.FindByXPath("/html/body/div[4]/div/div/p[2]/button"))
	browser.ClickElement(browser.FindByXPath("/html/body/div[5]/ul[1]/li[22]/button"))
	browser.ClickElement(browser.FindByXPath("/html/body/div[4]/div/div/label/input"))

	doNothing := func() error { return nil }
	browser.ExecuteWithWait(doNothing)

	html, _ := browser.HTML()

	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(html))
	return scrapeBattleLinks(doc)
}

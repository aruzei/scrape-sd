package showdown

import (
	"math"
	"scrape-sd/browser"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type BattleRoom struct {
	Browser *browser.Browser
}

func NewBattleRoom() *BattleRoom {
	browser := browser.CreateChrome()
	doNothing := func() error { return nil }
	browser.ExecuteWithWait(doNothing)

	browser.NavigatePage(battleULR)

	room := BattleRoom{
		Browser: &browser}

	return &room
}

// Scrape scrapes html in http://play.pokemonshowdown.com/battles
func (room *BattleRoom) Scrape() []BattleLink {

	browser := room.Browser

	browser.ClickElement(browser.FindByXPath(xpath_battle_Format))
	browser.ClickElement(browser.FindByButton(label_format_VGC2017))
	browser.ClickElement(browser.FindByXPath(xpath_EOLPlus1300))

	doNothing := func() error { return nil }
	browser.ExecuteWithWait(doNothing)

	html, _ := browser.HTML()

	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(html))
	return scrapeBattleLinks(doc)
}

func (room *BattleRoom) executeDownLoad(links []BattleLink, ch chan int) {
	// room := NewBattleRoom()

	for index := range links {
		room.Browser.NavigatePage(links[index].URL)
		room.Browser.WaitSeconds(30)
		room.Browser.ClickElement(room.Browser.FindByXPath(xpath_download_battle))
	}
	room.Browser.WaitSeconds(10)
	defer close(ch)
}

func divideLinks(links []BattleLink, division int) [][]BattleLink {
	var chunks [][]BattleLink
	sliceSize := len(links)
	size := int(math.Ceil(float64(sliceSize) / float64(division))) // 10/4 = 2.25 > 3

	for i := 0; i < sliceSize; i += size {
		end := i + size
		if sliceSize < end {
			end = sliceSize
		}
		chunks = append(chunks, links[i:end])
	}

	return chunks
}

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

	browser.ClickElement(browser.FindByXPath("/html/body/div[4]/div/div/p[2]/button"))
	browser.ClickElement(browser.FindByXPath("/html/body/div[5]/ul[1]/li[22]/button"))
	browser.ClickElement(browser.FindByXPath("/html/body/div[4]/div/div/label/input"))

	doNothing := func() error { return nil }
	browser.ExecuteWithWait(doNothing)

	html, _ := browser.HTML()

	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(html))
	return scrapeBattleLinks(doc)
}

// DownLoadBattles dowloads from linkes battle concurrently.
// Html file is saved at a defined by the browser.
func DownLoadBattles(links []BattleLink, division int) {

	executeDL := func(links []BattleLink,
		ch chan int) {
		downLoadBattles(links, ch)
	}
	splitLinks := divideLinks(links, division)
	channels := make([]chan int, division)
	for index := range splitLinks {
		channels[index] = make(chan int, 1)
		go executeDL(splitLinks[index], channels[index])
	}
	for ch := range channels {
		<-channels[ch]
	}
}

func downLoadBattles(links []BattleLink, ch chan int) {
	room := NewBattleRoom()

	for index := range links {
		room.Browser.NavigatePage(links[index].URL)
		room.Browser.WaitSeconds(120)
		room.Browser.ClickElement(room.Browser.FindByXPath("/html/body/div[4]/div[5]/div/p[1]/span/a"))
	}
	room.Browser.WaitSeconds(30)
	defer close(ch)
	defer room.Browser.Destroy()
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

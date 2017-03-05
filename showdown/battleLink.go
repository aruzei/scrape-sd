package showdown

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type BattleLink struct {
	URL     string
	Text    string
	IsValid bool
}

func scrapeBattleLinks(doc *goquery.Document) []BattleLink {
	links := make([]BattleLink, 0)
	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		link := newBattleLink(s)
		if link.IsValid {
			links = append(links, *link)
		}
	})
	return links
}

func newBattleLink(s *goquery.Selection) *BattleLink {
	url, _ := s.Attr("href")
	text := s.Text()
	return &BattleLink{
		URL:     rootURL + url,
		Text:    text,
		IsValid: strings.Contains(url, "/battle-gen7vgc2017")}
}

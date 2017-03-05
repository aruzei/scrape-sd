package main

import (
	"scrape-sd.local/showdown"
	"scrape-sd.local/test/showdown"
)

func main() {
	battleRoom := showdown.NewBattleRoom()
	defer battleRoom.Browser.Stop()

	links := battleRoom.Scrape()

	showdown.DownLoadBattles(links, 8)
	testshowdown.TestNewBattleResult("./.debug/dummy_battle1.html")
	testshowdown.TestNewBattleResult("./.debug/dummy_battle2.html")
}

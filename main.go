package main

import (
	"scrape-sd.local/showdown"
	"scrape-sd.local/util"
)

func main() {
	battleRoom := showdown.NewBattleRoom()
	defer battleRoom.Browser.Stop()

	links := battleRoom.Scrape()

	showdown.DownLoadBattles(links, 8)

	util.FlushAsCSV(links, "./.debug/links.csv")
	battleRoom.Browser.Screenshot("./.debug/debug.jpg")
}

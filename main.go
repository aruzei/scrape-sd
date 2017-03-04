package main

import "scrape-sd.local/showdown"

func main() {
	battleRoom := showdown.NewBattleRoom()
	defer battleRoom.Browser.Stop()

	links := battleRoom.Scrape()

	showdown.DownLoadBattles(links, 8)
}

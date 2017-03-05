package main

import (
	"scrape-sd/showdown"
	"scrape-sd/test/showdown"
)

func main() {
	battleRoom := showdown.NewBattleRoom()
	defer battleRoom.Browser.Stop()

	links := battleRoom.Scrape()

	showdown.DownLoadBattles(links, 4)

}
func test() {
	testshowdown.TestNewBattleResult()
}

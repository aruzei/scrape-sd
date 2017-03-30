package main

import (
	"io/ioutil"
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
func dumpResults(downLoadDir string) {
	fis, err := ioutil.ReadDir(downLoadDir)

	if err != nil {
		panic(err)
	}

	for fileIndex := range fis {
		result := showdown.NewResult(downLoadDir + fis[fileIndex].Name())
		result.Dump("./.debug/battle.csv")
	}

}

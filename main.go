package main

import (
	"io/ioutil"
	"scrape-sd/showdown"
	"scrape-sd/test/showdown"
)

func main() {
	getLinks := func() []showdown.BattleLink {
		battleRoom := showdown.NewBattleRoom()
		defer battleRoom.Browser.Stop()
		return battleRoom.Scrape()
	}
	downLoad := func(links []showdown.BattleLink, downLoader *showdown.DownLoader) {
		downLoader.DownLoadBattles(links)
	}

	downLoader := showdown.NewDownLoader(4)
	defer downLoader.DestroyBrowsers()

	for index := 0; index < 2; index++ {
		links := getLinks()
		downLoad(links, &downLoader)
	}

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

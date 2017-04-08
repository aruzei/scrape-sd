package showdown

import "math"

type DownLoader struct {
	Rooms []BattleRoom
}

// NewDownLoader return numbers of BattleRoom
func NewDownLoader(number int) DownLoader {
	rooms := make([]BattleRoom, number)

	for index := 0; index < len(rooms); index++ {
		room := *NewBattleRoom()
		rooms[index] = room
	}
	return DownLoader{Rooms: rooms}
}

// DownLoadBattles dowloads from linkes battle concurrently.
// Html file is saved at a defined by the browser.
func (downLoader *DownLoader) DownLoadBattles(links []BattleLink) {

	executeDL := func(room *BattleRoom, links []BattleLink,
		ch chan int) {
		room.executeDownLoad(links, ch)
	}
	division := len(downLoader.Rooms)
	splitLinks := divideLinks(links, division)
	totalChannels := int32(math.Min(float64(division), float64(len(links))))
	channels := make([]chan int, totalChannels)
	for index := range splitLinks {
		channels[index] = make(chan int, 1)
		go executeDL(&downLoader.Rooms[index], splitLinks[index], channels[index])
	}
	for ch := range channels {
		<-channels[ch]
	}
}

// DestroyBrowsers destroies inner browsers.
func (downLoader *DownLoader) DestroyBrowsers() {
	for index := 0; index < len(downLoader.Rooms); index++ {
		downLoader.Rooms[index].Browser.Destroy()
	}
}

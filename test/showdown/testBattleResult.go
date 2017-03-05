package testshowdown

import "scrape-sd.local/showdown"
import "fmt"

func TestNewBattleResult(html string) {
	result := showdown.NewResult(html)
	fmt.Println(result.Winner.ID)
	fmt.Println(result.Loser.ID)

	fmt.Println(result.Winner.Rate)
	fmt.Println(result.Loser.Rate)

	fmt.Println(result.Winner.Team)
	fmt.Println(result.Loser.Team)
}

package testshowdown

import "scrape-sd/showdown"
import "fmt"

// TestNewBattleResult verifies showdown.NewBattleResult.
// Case1 dummy_battle1.html
// Case2 dummy_battle2.html
func TestNewBattleResult() {
	result1 := showdown.Result{
		Winner: showdown.Player{
			ID:   "Elenazi",
			Rate: "1485",
			Team: showdown.Team{
				ID:       "p1",
				Pokemons: []string{"Arcanine", "Garchomp", "Muk-Alola", "Porygon2", "Tapu Bulu", "Tapu Koko"},
			},
		},
		Loser: showdown.Player{
			ID:   "TTshen",
			Rate: "1436",
			Team: showdown.Team{
				ID:       "p2",
				Pokemons: []string{"Lilligant", "Oranguru", "Oricorio-Pom-Pom", "Raichu-Alola", "Tapu Koko", "Torkoal"},
			},
		},
	}
	result2 := showdown.Result{
		Winner: showdown.Player{
			ID:   "ChoiceSpecsFakeOut",
			Rate: "1613",
			Team: showdown.Team{
				ID:       "p1",
				Pokemons: []string{"Arcanine", "Garchomp", "Gengar", "Mandibuzz", "Ninetales-Alola", "Tapu Lele"},
			},
		},
		Loser: showdown.Player{
			ID:   "Kurono_kei",
			Rate: "1641",
			Team: showdown.Team{
				ID:       "p2",
				Pokemons: []string{"Arcanine", "Celesteela", "Dhelmise", "Garchomp", "Tapu Fini", "Tapu Koko"},
			},
		},
	}
	fmt.Println(testNewBattleResult("./test/showdown/dummy_battle1.html", result1))
	fmt.Println(testNewBattleResult("./test/showdown/dummy_battle2.html", result2))
}

func testNewBattleResult(html string, expected showdown.Result) bool {

	result := showdown.NewResult(html)
	return isPlayerSame(result.Winner, expected.Winner) &&
		isPlayerSame(result.Loser, expected.Loser)
}
func isPlayerSame(player1 showdown.Player, player2 showdown.Player) bool {
	if !(player1.ID == player2.ID &&
		player1.Rate == player2.Rate &&
		player1.Team.ID == player2.Team.ID) {
		return false
	}
	if !(len(player1.Team.Pokemons) == (len(player2.Team.Pokemons))) {
		return false
	}
	for index := 0; index < len(player1.Team.Pokemons); index++ {
		if player1.Team.Pokemons[index] != player2.Team.Pokemons[index] {
			return false
		}
	}
	return true
}

package showdown

import (
	"io/ioutil"
	"sort"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/microcosm-cc/bluemonday"
)

type Result struct {
	Winner Player
	Loser  Player
}

type Player struct {
	ID   string
	Rate string
	Team Team
}
type Team struct {
	ID       string
	Pokemons []string
}

func NewResult(html string) Result {
	log := scrapeBattleLog(html)
	teams := log.createTeams()
	players := log.createPlayers(teams)
	return Result{Winner: players[0], Loser: players[1]}
}

func newPlayer(id string, rate string, team Team) Player {
	return Player{ID: id, Rate: rate, Team: team}
}

func newTeam(id string, pokemons []string) Team {
	return Team{ID: id, Pokemons: pokemons}
}

type battleLog string

func (log battleLog) createTeams() []Team {
	var pkmns1 []string
	var pkmns2 []string
	ans := strings.Split(string(log), "\n")
	for index := 0; index < len(ans); index++ {
		if strings.Contains(ans[index], "|poke|p") {
			if strings.Contains(ans[index], "|poke|p1") {
				pkmn := strings.Split(strings.Split(ans[index], "|poke|p1|")[1], ",")[0]
				pkmns1 = append(pkmns1, pkmn)
			} else {
				pkmn := strings.Split(strings.Split(ans[index], "|poke|p2|")[1], ",")[0]
				pkmns2 = append(pkmns2, pkmn)
			}
		}
	}
	var teams []Team
	sort.Strings(pkmns1)
	sort.Strings(pkmns2)

	teams = append(teams, newTeam("p1", pkmns1))
	teams = append(teams, newTeam("p2", pkmns2))

	return teams
}
func (log battleLog) createPlayers(teams []Team) []Player {
	var player1, player2 string
	var rate1, rate2 string

	ans := strings.Split(string(log), "\n")
	for index := 0; index < len(ans); index++ {
		if strings.Contains(ans[index], "|player|p") {
			if strings.Contains(ans[index], "|player|p1|") {
				player1 = strings.Split(strings.Split(ans[index], "|player|p1|")[1], "|")[0]
			} else if strings.Contains(ans[index], "|player|p2|") {
				player2 = strings.Split(strings.Split(ans[index], "|player|p2|")[1], "|")[0]
			}
		} else if strings.Contains(ans[index], "|raw|") {
			if strings.Contains(ans[index], player1) {
				rate1 = strings.Split(strings.Split(ans[index], "<strong>")[1], "<\\/strong>")[0]
			} else {
				rate2 = strings.Split(strings.Split(ans[index], "<strong>")[1], "<\\/strong>")[0]
			}
		}
	}
	var players []Player
	players = append(players, newPlayer(player1, rate1, teams[0]))
	players = append(players, newPlayer(player2, rate2, teams[1]))

	return players
}
func (log battleLog) createResult(players []Player) Result {
	var winner, loser Player
	ans := strings.Split(string(log), "\n")
	for index := 0; index < len(ans); index++ {
		if strings.Contains(ans[index], "|win|") {
			if strings.Contains(ans[index], players[0].ID) {
				winner = players[0]
				loser = players[1]
			} else {
				winner = players[1]
				loser = players[0]
			}
		}
	}
	return Result{Winner: winner, Loser: loser}
}
func scrapeBattleLog(html string) battleLog {
	fileInfos, _ := ioutil.ReadFile(html)
	stringReader := strings.NewReader(string(fileInfos))
	doc, _ := goquery.NewDocumentFromReader(stringReader)
	res, _ := doc.Find("body").Html()

	p := bluemonday.NewPolicy()
	p.AllowElements("script")
	return battleLog(p.Sanitize(res))
}

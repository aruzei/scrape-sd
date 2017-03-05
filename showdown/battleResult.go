package showdown

type Result struct {
	Winner Player
	Loser  Player
}

type Player struct {
	ID   string
	Rate int
	Team Team
}
type Team []string

func NewResult(html string) Result {
	return Result{}
}

func NewPlayer(ID string, rate int, team Team) Player {
	return Player{}
}

func NewTeam(pokemons string) Team {
	return Team{}
}

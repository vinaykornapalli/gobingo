package model

type GameState struct {
	crntActivePlayer int
	chosenNumber     int
	maxPlayers       int
}

func (s *GameState) updateState() {

	if s.crntActivePlayer == s.maxPlayers-1 {
		s.crntActivePlayer = 0
	} else {
		s.crntActivePlayer = s.crntActivePlayer + 1
	}
}

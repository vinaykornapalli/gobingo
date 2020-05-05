package model

type GameState struct {
	crntActivePlayer int
	chosenNumber     int
	maxPlayers       int
	finishedNumbers  []int
}

func (s *GameState) UpdateState() {

	s.finishedNumbers = append(s.finishedNumbers, s.chosenNumber)
	s.chosenNumber = 0
	if s.crntActivePlayer == s.maxPlayers-1 {
		s.crntActivePlayer = 0
	} else {
		s.crntActivePlayer = s.crntActivePlayer + 1
	}
}

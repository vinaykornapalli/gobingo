package model

type GameState struct {
	CrntActivePlayer int
	ChosenNumber     int
	MaxPlayers       int
	FinishedNumbers  []int
}

func (s *GameState) UpdateState() {

	s.FinishedNumbers = append(s.FinishedNumbers, s.ChosenNumber)
	if s.CrntActivePlayer == s.MaxPlayers-1 {
		s.CrntActivePlayer = 0
	} else {
		s.CrntActivePlayer = s.CrntActivePlayer + 1
	}
}

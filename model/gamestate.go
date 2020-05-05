package model

//GameState structure
type GameState struct {
	CrntActivePlayer int  `json:"active_player"`
	ChosenNumber     int   `json:"chosen_number"`
	MaxPlayers       int   `json:"max_players"`
	FinishedNumbers  []int  `json:"finished_numbers"`
}

//UpdateState updates the crnt state
func (s *GameState) UpdateState() {
	s.FinishedNumbers = append(s.FinishedNumbers, s.ChosenNumber)
	s.ChosenNumber = 0
	if s.CrntActivePlayer == s.MaxPlayers-1 {
		s.CrntActivePlayer = 0
	} else {
		s.CrntActivePlayer = s.CrntActivePlayer + 1
	}
}

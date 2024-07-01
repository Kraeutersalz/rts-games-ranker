package data

type Player struct {
	Name   string
	Wins   int
	Losses int
	Rating float64
}

func TotalMatches(p Player) int {
	return p.Wins + p.Losses
}

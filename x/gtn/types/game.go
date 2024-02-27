package types

func (g Game) IsReveald() bool {
	return g.Number != 0
}

func (g Game) SplitWinnersAndLosers(guesses []Guess, dist uint64) (winners []Guess, losers []Guess) {
	for _, guess := range guesses {
		// TODO: Consider the underflow and overflow cases
		if isInBetween(guess.Number, g.Number-dist, g.Number+dist) {
			winners = append(winners, guess)
		} else {
			losers = append(losers, guess)
		}
	}
	return
}

func isInBetween(n, min, max uint64) bool {
	return n >= min && n <= max
}

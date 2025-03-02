package faircointoss

import "math/rand/v2"

func UnfairCoinToss() CoinSide {
	randomNumber := rand.Float64()

	if randomNumber < 0.99 {
		return CoinSideHeads
	} else {
		return CoinSideTails
	}
}

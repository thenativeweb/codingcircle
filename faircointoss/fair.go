package faircointoss

func FairCoinToss() CoinSide {
	for {
		firstCoinSide := UnfairCoinToss()
		secondCoinSide := UnfairCoinToss()

		if firstCoinSide == secondCoinSide {
			continue
		}

		return firstCoinSide
	}
}

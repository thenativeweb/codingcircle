package markovcount

import "math/rand"

type Transition struct {
	From        string
	To          string
	Probability float64
}

func Count(
	initialState string,
	transitions []Transition,
	iterations int,
) map[string]int {
	transitionResults := map[string]int{}

	currentState := initialState
	for i := 0; i < iterations; i++ {
		nextState := getNextState(currentState, transitions)

		if _, ok := transitionResults[nextState]; !ok {
			transitionResults[nextState] = 0
		}
		transitionResults[nextState]++

		currentState = nextState
	}

	return transitionResults
}

func getNextState(
	currentState string,
	transitions []Transition,
) string {
	randomValue := rand.Float64()

	for _, transition := range transitions {
		if transition.From != currentState {
			continue
		}

		if randomValue <= transition.Probability {
			return transition.To
		}

		randomValue -= transition.Probability
	}

	panic("no transition found")
}

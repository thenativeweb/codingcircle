package busfactor

import (
	"errors"
	"time"
)

type Event struct {
	Time  time.Time
	Type  string
	Count int
}

type Result struct {
	Count int
	From  time.Time
	Until time.Time
}

func GetMaximumLoad(events []Event) Result {
	currentLoad := 0
	maxLoad := 0
	maxLoadStartTime := events[0].Time
	maxLoadEndTime := events[0].Time

	for i, event := range events {
		switch event.Type {
		case "joined":
			currentLoad += event.Count
		case "left":
			currentLoad -= event.Count
		default:
			panic(errors.New("unknown event type"))
		}

		if currentLoad > maxLoad {
			maxLoad = currentLoad
			maxLoadStartTime = event.Time
			if i+1 < len(events) {
				maxLoadEndTime = events[i+1].Time
			} else {
				maxLoadEndTime = event.Time
			}
		}
	}

	return Result{
		Count: maxLoad,
		From:  maxLoadStartTime,
		Until: maxLoadEndTime,
	}
}

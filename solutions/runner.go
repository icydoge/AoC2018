package solutions

import (
	"time"

	"github.com/icydoge/AoC2018/solutions/day1"
)

var runMap = map[string]interface{}{
	"1": day1.Run,
}

func RunSolution(solution string, input string) (string, time.Duration) {
	var results string
	start := time.Now()

	if f, ok := runMap[solution]; ok {
		results = f.(func(string) string)(input)
	} else {
		results = "(Solution not found)"
	}
	realExecutionTime := time.Now().Sub(start)
	return results, realExecutionTime
}

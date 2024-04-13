package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	rando "github.com/jadefish/cata-rando"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s member1 member2 ...\n", os.Args[0])
		os.Exit(1)
	}

	peeps := os.Args[1:]
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(peeps), func(i, j int) { peeps[i], peeps[j] = peeps[j], peeps[i] })

	r := rando.NewRoller()
	picks := make([]rando.Pick, 0, len(peeps))

	// Pick a tank and a healer spec:
	picks = append(picks, r.Roll(rando.Tank))
	picks = append(picks, r.Roll(rando.Healer))

	// then fill with DPS:
	n := max(0, len(peeps)-len(picks))

	for i := 0; i < n; i++ {
		picks = append(picks, r.Roll(rando.DPS))
	}

	assignments := make(map[string]rando.Pick, len(peeps))

	for i, peep := range peeps {
		assignments[peep] = picks[i]
	}

	for peep, assignment := range assignments {
		fmt.Printf("%s: %s\n", peep, assignment)
	}
}

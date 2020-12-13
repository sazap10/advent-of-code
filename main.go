package main

import (
	"log"
	"time"

	"github.com/bugsnag/bugsnag-go"
	"github.com/pkg/errors"
	"github.com/sazap10/advent-of-code/pkg/solution"
	"github.com/sazap10/advent-of-code/solutions"
)

func main() {
	bugsnag.Configure(bugsnag.Configuration{
		// Your Bugsnag project API key
		APIKey: "4833332d828180db8e91c20ce3b85e74",
		// The import paths for the Go packages containing your source files
		ProjectPackages: []string{"main", "github.com/sazap10/advent-of-code"},
	})

	solutions2020 := solutions.NewSolutions2020()

	runSolutions(solutions2020, "2020")
}

func runSolutions(solutionsMap solution.Map, year string) {
	log.Printf("Year %s solutions:\n", year)
	for _, s := range solutionsMap {
		log.Printf("Running solution: %s, URL: %s\n", s.Name(), s.URL())

		beforeRunTime := time.Now()
		result, err := s.Run()
		afterRunTime := time.Now()
		timeDiff := afterRunTime.Sub(beforeRunTime).Milliseconds()
		if err != nil {
			bugsnag.Notify(errors.Wrapf(err, "error running solution '%s'", s.Name()))
		}

		log.Printf("Result: %s, took %d ms\n", result, timeDiff)
	}
}

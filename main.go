package main

import (
	"log"
	"time"

	"github.com/bugsnag/bugsnag-go"
	"github.com/pkg/errors"
	"github.com/sazap10/advent-of-code/solutions"
)

func main() {
	bugsnag.Configure(bugsnag.Configuration{
		// Your Bugsnag project API key
		APIKey: "4833332d828180db8e91c20ce3b85e74",
		// The import paths for the Go packages containing your source files
		ProjectPackages: []string{"main", "github.com/sazap10/advent-of-code"},
	})

	solutions := solutions.NewSolutions2020()

	for _, s := range solutions {
		log.Printf("Running solution: %s, URL: %s\n", s.Name(), s.URL())

		beforeRunTime := time.Now()
		err, result := s.Run()
		afterRunTime := time.Now()
		timeDiff := afterRunTime.Sub(beforeRunTime).Milliseconds()
		if err != nil {
			bugsnag.Notify(errors.Wrapf(err, "error running solution '%s'", s.Name()))
		}

		log.Printf("Result: %s, took %d ms\n", result, timeDiff)
	}
}

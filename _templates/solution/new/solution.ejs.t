---
to: solutions/<%= year %>/solution<%= day %>/solution.go
---
package solution<%= day %>

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
	"github.com/sazap10/advent-of-code/pkg/solution"
)

// Solution holds any data for the program
type Solution struct {
	solution.Label
}

// New instantiates the solution
func New() *Solution {
	label := solution.NewLabel("<%= name %>", "https://adventofcode.com/<%= year %>/day/<%= day %>", "<%= year %>")
	return &Solution{
		Label: label,
	}
}

// Run contains the logic to solving the problem
func (s *Solution) Run() (string, error) {
	sampleInput, err := s.getInput("solutions/<%= year %>/solution<%= day %>/sample.txt")
	if err != nil {
		return "", errors.Wrapf(err, "Unable to read sample file")
	}

	input, err := s.getInput("solutions/<%= year %>/solution<%= day %>/input.txt")
	if err != nil {
		return "", errors.Wrapf(err, "Unable to read input file")
	}

	part1 := s.part1(sampleInput)

	part2 := s.part2(sampleInput)

	log.Printf("Sample answers; Part 1: %s, Part 2: %s", part1, part2)

	part1 = s.part1(input)

	part2 = s.part2(input)

	return fmt.Sprintf("Part 1: %s, Part 2: %s", part1, part2), nil
}

func (s *Solution) getInput(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	output := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		a := scanner.Text()

		output = append(output, a)
	}
	return output, scanner.Err()
}

func (s *Solution) part1(input []string) string {
	answer := 0

	return fmt.Sprint(answer)
}

func (s *Solution) part2(input []string) string {
	answer := 0

	return fmt.Sprint(answer)
}



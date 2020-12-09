package solution6

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
	label := solution.NewLabel("Custom Customs", "https://adventofcode.com/2020/day/6", "2020")
	return &Solution{
		Label: label,
	}
}

// Run contains the logic to solving the problem
func (s *Solution) Run() (string, error) {
	sampleInput, err := s.getInput("solutions/2020/solution6/sample.txt")
	if err != nil {
		return "", errors.Wrapf(err, "Unable to read sample file")
	}

	input, err := s.getInput("solutions/2020/solution6/input.txt")
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

func (s *Solution) getInput(path string) ([]map[rune]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var output []map[rune]int
	block := make(map[rune]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()

		if len(l) != 0 {
			for _, c := range l {
				block[c]++
			}
			continue
		}

		if len(l) == 0 {
			output = append(output, block)
			block = make(map[rune]int)
			continue
		}
	}
	if len(block) > 0 {
		output = append(output, block)
	}
	return output, scanner.Err()
}

func (s *Solution) part1(input []map[rune]int) string {
	sumCounts := sumAnyQuestionAnsweredYes(input)

	return fmt.Sprint(sumCounts)
}

func (s *Solution) part2(input []map[rune]int) string {
	sumCounts := sumEveryQuestionAnsweredYes(input)

	return fmt.Sprint(sumCounts)
}

// sumAnyQuestionAnsweredYes calculates the sum of anyone answering
// a question with a yes
func sumAnyQuestionAnsweredYes(forms []map[rune]int) int {
	total := 0
	for _, f := range forms {
		total += len(f)
	}
	return total
}

// sumEveryQuestionAnsweredYes calculates the sum of everyone answering
// a question with a yes
func sumEveryQuestionAnsweredYes(forms []map[rune]int) int {
	total := 0
	for _, f := range forms {
		total += len(f)
	}
	return total
}

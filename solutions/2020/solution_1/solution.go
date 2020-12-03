package solution_1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/pkg/errors"
	"github.com/sazap10/advent-of-code/pkg/solution"
)

// Solution holds any data for the program
type Solution struct {
	solution.Label
	Sum int
}

// New instantiates the solution
func New() *Solution {
	label := solution.NewLabel("Multiply numbers which sum to 2020", "https://adventofcode.com/2020/day/1", "2020")
	return &Solution{
		Label: label,
		Sum: 2020,
	}
}

// Run contains the logic to solving the problem
func (s *Solution) Run() (string, error) {
	numbers, err := s.getInput()

	if err != nil {
		return "", errors.Wrapf(err, "Unable to read input file")
	}

	part1 := s.part1(&numbers)

	part2 := s.part2(&numbers)

	return fmt.Sprintf("Part 1: %s, Part 2: %s", part1, part2), nil
}

func (s *Solution) getInput() ([]int, error) {
	inputFile := "solutions/2020/solution_1/input.txt"
	file, err := os.Open(inputFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
				return lines, err
		}
		lines = append(lines, x)
	}
	return lines, scanner.Err()
}

func (s *Solution) part1(numbers *[]int) string{
	for index, i := range *numbers {
		numbersWithoutI := remove(*numbers, index)
		for _, j := range numbersWithoutI {
			if i + j == s.Sum {
				total := i*j
				return fmt.Sprint(total)
			}
		}
	}
	return "Could not find answer"
}

func (s *Solution) part2(numbers *[]int) string{
	log.Println(len(*numbers))
	for index, i := range *numbers {
		numbersWithoutI := remove(*numbers, index)
		for index2, j := range numbersWithoutI {
			numbersWithoutJ := remove(numbersWithoutI, index2)
			for _, k := range numbersWithoutJ {
				if (i + j + k) == s.Sum {
					total := i*j*k
					return fmt.Sprint(total)
				}
			}
		}
	}
	return "Could not find answer"
}

func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	// We do not need to put s[i] at the end, as it will be discarded anyway
	return s[:len(s)-1]
}
package solution9

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/pkg/errors"
	"github.com/sazap10/advent-of-code/pkg/solution"
)

// Solution holds any data for the program
type Solution struct {
	solution.Label
}

// New instantiates the solution
func New() *Solution {
	label := solution.NewLabel("Encoding Error", "https://adventofcode.com/2020/day/9", "2020")
	return &Solution{
		Label: label,
	}
}

// Run contains the logic to solving the problem
func (s *Solution) Run() (string, error) {
	sampleInput, err := s.getInput("solutions/2020/solution9/sample.txt")
	if err != nil {
		return "", errors.Wrapf(err, "Unable to read sample file")
	}

	input, err := s.getInput("solutions/2020/solution9/input.txt")
	if err != nil {
		return "", errors.Wrapf(err, "Unable to read input file")
	}

	part1 := s.part1(sampleInput, 5)

	part2 := s.part2(sampleInput, 127)

	log.Printf("Sample answers; Part 1: %s, Part 2: %s", part1, part2)

	part1 = s.part1(input, 25)

	part2 = s.part2(input, 18272118)

	return fmt.Sprintf("Part 1: %s, Part 2: %s", part1, part2), nil
}

func (s *Solution) getInput(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	output := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		a, err := strconv.Atoi(scanner.Text())

		if err != nil {
			return output, errors.New(fmt.Sprintf("Unable to convert: %s", scanner.Text()))
		}

		output = append(output, a)
	}
	return output, scanner.Err()
}

func (s *Solution) part1(input []int, preambleLength int) string {
	answer := 0

	for i := preambleLength; i < len(input); i++ {
		slice := input[i-preambleLength : i]
		if !isSumValid(slice, input[i]) {
			answer = input[i]
			break
		}
	}

	return fmt.Sprint(answer)
}

func (s *Solution) part2(input []int, targetNumber int) string {
	list := findListSum(input, targetNumber)

	sort.Ints(list)

	answer := list[0] + list[len(list)-1]

	return fmt.Sprint(answer)
}

func isSumValid(input []int, targetNumber int) bool {
	valid := false

	for i, a := range input {
		for j, b := range input {
			if i != j && (a+b) == targetNumber {
				valid = true
			}
		}
	}

	return valid
}

func findListSum(input []int, targetNumber int) []int {
	currentStart := 0
	currentIndex := 0
	currentList := []int{}

	for {
		currentList = append(currentList, input[currentIndex])
		sumArray := sum(currentList)
		switch {
		case sumArray == targetNumber:
			return currentList
		case sumArray < targetNumber:
			currentIndex++
		default:
			currentStart++
			currentIndex = currentStart
			currentList = []int{}
		}
	}
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

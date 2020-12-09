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

// CustomForm holds the data of the yes answer to questions on the custom form
type CustomForm []rune

// CustomFormGroup holds the data for a group of CustomForms
type CustomFormGroup []CustomForm

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

func (s *Solution) getInput(path string) ([]CustomFormGroup, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var output []CustomFormGroup
	block := CustomFormGroup{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()

		if len(l) != 0 {
			f := CustomForm{}
			for _, c := range l {
				f = append(f, c)
			}
			block = append(block, f)
			continue
		}

		if len(l) == 0 {
			output = append(output, block)
			block =  CustomFormGroup{}
			continue
		}
	}
	if len(block) > 0 {
		output = append(output, block)
	}
	return output, scanner.Err()
}

func (s *Solution) part1(input []CustomFormGroup) string {
	sumCounts := sumAnyQuestionAnsweredYes(input)

	return fmt.Sprint(sumCounts)
}

func (s *Solution) part2(input []CustomFormGroup) string {
	sumCounts := sumEveryQuestionAnsweredYes(input)

	return fmt.Sprint(sumCounts)
}

// sumAnyQuestionAnsweredYes calculates the sum of anyone answering
// a question with a yes
func sumAnyQuestionAnsweredYes(formGroups []CustomFormGroup) int {
	total := 0
	for _, fg := range formGroups {
		answers := make(map[rune]int)
		for _, f := range fg {
			for _, r := range f {
				answers[r]++
			}
		}
		total += len(answers)
	}
	return total
}

// sumEveryQuestionAnsweredYes calculates the sum of everyone answering
// a question with a yes
func sumEveryQuestionAnsweredYes(formGroups []CustomFormGroup) int {
	total := 0
	for _, fg := range formGroups {
		answers := []rune{}
		for i, f := range fg {
			if i == 0 {
				answers = f
			} else {
				answers = intersect(answers, f)
			}
		}
		total += len(answers)
	}
	return total
}

func intersect(a, b []rune) (c []rune){
	m := make(map[rune]bool)

	for _, item := range a {
		m[item] = true
	}
	
	for _, item := range b {
		if _, ok := m[item]; ok {
			c = append(c, item)
		}
	}

	return
}
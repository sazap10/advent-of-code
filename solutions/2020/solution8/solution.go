package solution8

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/sazap10/advent-of-code/pkg/solution"
)

// Solution holds any data for the program
type Solution struct {
	solution.Label
}

type Instruction struct {
	Command  string
	Argument int
}

// New instantiates the solution
func New() *Solution {
	label := solution.NewLabel("Handheld Halting", "https://adventofcode.com/2020/day/8", "2020")
	return &Solution{
		Label: label,
	}
}

// Run contains the logic to solving the problem
func (s *Solution) Run() (string, error) {
	sampleInput, err := s.getInput("solutions/2020/solution8/sample.txt")
	if err != nil {
		return "", errors.Wrapf(err, "Unable to read sample file")
	}

	input, err := s.getInput("solutions/2020/solution8/input.txt")
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

func (s *Solution) getInput(path string) ([]Instruction, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	output := []Instruction{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		splitLine := strings.Split(scanner.Text(), " ")

		i := Instruction{
			Command: splitLine[0],
		}

		a, err := strconv.Atoi(splitLine[1])

		if err != nil {
			return output, errors.New(fmt.Sprintf("Unable to convert argument: %s for line %s", splitLine[1], scanner.Text()))
		}

		i.Argument = a

		output = append(output, i)
	}
	return output, scanner.Err()
}

func (s *Solution) part1(input []Instruction) string {
	acc, _ := runInstructions(input)

	return fmt.Sprint(acc)
}

func (s *Solution) part2(input []Instruction) string {
	acc := 0

	for i, instruction := range input {
		switch instruction.Command {
		case "nop":
			newInstruction := Instruction{Command: "jmp", Argument: instruction.Argument}
			if a, loop := testChangeInstructions(i, newInstruction, input); !loop {
				acc = a
				break
			}
		case "jmp":
			newInstruction := Instruction{Command: "nop", Argument: instruction.Argument}
			if a, loop := testChangeInstructions(i, newInstruction, input); !loop {
				acc = a
				break
			}
		}
	}

	return fmt.Sprint(acc)
}

func testChangeInstructions(index int, newInstruction Instruction, input []Instruction) (int, bool) {
	newInstructions := make([]Instruction, len(input))
	copy(newInstructions, input)
	newInstructions[index] = newInstruction
	return runInstructions(newInstructions)
}

func runInstructions(input []Instruction) (int, bool) {
	acc := 0
	instructionToRun := 0
	instructionList := make(map[int]bool)
	for {
		if instructionList[instructionToRun] {
			return acc, true
		}

		i := input[instructionToRun]
		instructionList[instructionToRun] = true

		switch i.Command {
		case "nop":
			instructionToRun++
		case "acc":
			acc += i.Argument
			instructionToRun++
		case "jmp":
			instructionToRun += i.Argument
		}

		if instructionToRun == len(input) {
			return acc, false
		}
	}
}

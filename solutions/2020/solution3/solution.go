package solution3

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
	TreeCharacter byte
	OpenSpaceCharacter byte
}

// Input holds data parsed from the input
type Input struct {
	PatternLength int
	PatternBytes [][]byte
}

// New instantiates the solution
func New() *Solution {
	label := solution.NewLabel("Toboggan Trajectory", "https://adventofcode.com/2020/day/3", "2020")
	return &Solution{
		Label: label,
		TreeCharacter: '#',
		OpenSpaceCharacter: '.',
	}
}

// Run contains the logic to solving the problem
func (s *Solution) Run() (string, error) {
	sampleInput, err := s.getInput("solutions/2020/solution3/sample.txt")
	input, err := s.getInput("solutions/2020/solution3/input.txt")

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

func (s *Solution) getInput(path string) (Input, error) {
	file, err := os.Open(path)
	if err != nil {
		return Input{}, err
	}
	defer file.Close()

	output := Input{
		PatternBytes: [][]byte{},
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineBytes := []byte(scanner.Text())
		output.PatternLength = len(lineBytes)
		output.PatternBytes = append(output.PatternBytes, lineBytes)
	}
	return output, scanner.Err()
}

func (s *Solution) part1(input Input) string{
	trees := treesEncountered(input, 3, 1, s.TreeCharacter)

	return fmt.Sprint(trees)
}

func (s *Solution) part2(input Input) string{
	trees := treesEncountered(input, 1, 1, s.TreeCharacter)
	trees *= treesEncountered(input, 3, 1, s.TreeCharacter)
	trees *= treesEncountered(input, 5, 1, s.TreeCharacter)
	trees *= treesEncountered(input, 7, 1, s.TreeCharacter)
	trees *= treesEncountered(input, 1, 2, s.TreeCharacter)
	return fmt.Sprint(trees)
}

func traverse(input [][]byte, startingX, startingY, patternLength, movementX, movementY int) (byte, int, int) {
	positionX := (startingX + movementX) % patternLength
	positionY := (startingY + movementY)
	return input[positionY][positionX],positionX, positionY
}

func treesEncountered(input Input, movementX, movementY int, treeCharacter byte) int {
	trees := 0
	currentX := 0
	currentY := 0
	length := len(input.PatternBytes)
	for currentY < length-1 {
		endLocation, endX, endY := traverse(input.PatternBytes, currentX, currentY, input.PatternLength, movementX, movementY)
		if endLocation == treeCharacter {
			trees++
		}
		currentX = endX
		currentY = endY
	}
	return trees
}
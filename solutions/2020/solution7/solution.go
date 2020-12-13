package solution7

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/sazap10/advent-of-code/pkg/solution"
)

// Solution holds any data for the program
type Solution struct {
	solution.Label
}

type BagKey struct {
	Adjective string
	Colour    string
}

type BagValue struct {
	Key    BagKey
	Number int
}

// New instantiates the solution
func New() *Solution {
	label := solution.NewLabel("Handy Haversacks", "https://adventofcode.com/2020/day/7", "2020")
	return &Solution{
		Label: label,
	}
}

// Run contains the logic to solving the problem
func (s *Solution) Run() (string, error) {
	sampleInput, err := s.getInput("solutions/2020/solution7/sample.txt")
	if err != nil {
		return "", errors.Wrapf(err, "Unable to read sample file")
	}

	// input, err := s.getInput("solutions/2020/solution7/input.txt")
	// if err != nil {
	// 	return "", errors.Wrapf(err, "Unable to read input file")
	// }

	log.Println(sampleInput)

	part1 := s.part1(sampleInput)

	part2 := s.part2(sampleInput)

	log.Printf("Sample answers; Part 1: %s, Part 2: %s", part1, part2)

	// part1 = s.part1(input)

	// part2 = s.part2(input)

	return fmt.Sprintf("Part 1: %s, Part 2: %s", part1, part2), nil
}

func (s *Solution) getInput(path string) (map[BagKey][]BagValue, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var lineRe = regexp.MustCompile(`([a-z]+) ([a-z]+) bags contain (.*)\.`)
	var containRe = regexp.MustCompile(`([1-9][0-9]*) ([a-z]+) ([a-z]+) bags`)
	output := make(map[BagKey][]BagValue)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		matches := lineRe.FindStringSubmatch(scanner.Text())

		key := BagKey{
			Adjective: matches[1],
			Colour:    matches[2],
		}

		containingBagsString := matches[3]

		containingBagsList := strings.Split(containingBagsString, ",")

		containingBags := []BagValue{}
		for _, b := range containingBagsList {
			if containRe.MatchString(b) {
				bagMatch := containRe.FindStringSubmatch(b)
				log.Println(bagMatch)
				number, err := strconv.Atoi(bagMatch[1])
				if err != nil {
					return output, err
				}
				bagValue := BagValue{
					Number: number,
					Key: BagKey{
						Adjective: bagMatch[2],
						Colour:    bagMatch[3],
					},
				}

				containingBags = append(containingBags, bagValue)
			}

		}

		output[key] = containingBags

	}
	return output, scanner.Err()
}

func (s *Solution) part1(input map[BagKey][]BagValue) string {
	sumCounts := 0

	return fmt.Sprint(sumCounts)
}

func (s *Solution) part2(input map[BagKey][]BagValue) string {
	sumCounts := 0

	return fmt.Sprint(sumCounts)
}

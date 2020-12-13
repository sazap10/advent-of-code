package solution2

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

// PasswordPolicy describes the data format for the input file lines
type PasswordPolicy struct {
	Number1  int
	Number2  int
	Letter   byte
	Password string
}

// New instantiates the solution
func New() *Solution {
	label := solution.NewLabel("Password Philosophy", "https://adventofcode.com/2020/day/2", "2020")
	return &Solution{
		Label: label,
	}
}

// Run contains the logic to solving the problem
func (s *Solution) Run() (string, error) {
	sampleList, err := s.getInput("solutions/2020/solution2/sample.txt")

	if err != nil {
		return "", errors.Wrapf(err, "Unable to read sample input file")
	}

	passwordList, err := s.getInput("solutions/2020/solution2/input.txt")

	if err != nil {
		return "", errors.Wrapf(err, "Unable to read input file")
	}

	part1 := s.part1(&sampleList)

	part2 := s.part2(&sampleList)

	log.Printf("Sample answers; Part 1: %s, Part 2: %s", part1, part2)

	part1 = s.part1(&passwordList)

	part2 = s.part2(&passwordList)

	return fmt.Sprintf("Part 1: %s, Part 2: %s", part1, part2), nil
}

func (s *Solution) getInput(path string) ([]PasswordPolicy, error) {
	var re = regexp.MustCompile(`([1-9]\d*)-([1-9]\d*) ([a-z]): ([a-z]+)`)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var output []PasswordPolicy
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		matches := re.FindStringSubmatch(scanner.Text())
		if len(matches) == 5 {
			number1, err := strconv.Atoi(matches[1])
			if err != nil {
				return output, err
			}
			number2, err := strconv.Atoi(matches[2])
			if err != nil {
				return output, err
			}
			letterBytes := []byte(matches[3])
			p := PasswordPolicy{
				Number1:  number1,
				Number2:  number2,
				Letter:   letterBytes[0],
				Password: matches[4],
			}
			output = append(output, p)
		} else {
			return output, errors.New("Unable to parse data")
		}
	}
	return output, scanner.Err()
}

func (s *Solution) part1(passwords *[]PasswordPolicy) string {
	valid := 0
	for _, p := range *passwords {
		count := strings.Count(p.Password, string(p.Letter))
		if count >= p.Number1 && count <= p.Number2 {
			valid++
		}
	}
	return fmt.Sprint(valid)
}

func (s *Solution) part2(passwords *[]PasswordPolicy) string {
	valid := 0
	for _, p := range *passwords {
		if (p.Password[p.Number1-1] == p.Letter) != (p.Password[p.Number2-1] == p.Letter) {
			valid++
		}
	}
	return fmt.Sprint(valid)
}

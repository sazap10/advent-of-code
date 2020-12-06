package solution4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
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

// Passport holds data parsed from the input
type Passport struct {
	BirthYear      int    `passport:"byr"`
	IssueYear      int    `passport:"iyr"`
	ExpirationYear int    `passport:"eyr"`
	Height         string `passport:"hgt"`
	HairColour     string `passport:"hcl"`
	EyeColour      string `passport:"ecl"`
	PassportID     string `passport:"pid"`
	CountryID      *int   `passport:"cid"`
}

// New instantiates the solution
func New() *Solution {
	label := solution.NewLabel("Passport Processing", "https://adventofcode.com/2020/day/4", "2020")
	return &Solution{
		Label: label,
	}
}

// Run contains the logic to solving the problem
func (s *Solution) Run() (string, error) {
	sampleInput, err := s.getInput("solutions/2020/solution4/sample.txt")
	input, err := s.getInput("solutions/2020/solution4/input.txt")

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

func (s *Solution) getInput(path string) ([]Passport, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var output []Passport
	block := Passport{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()

		if len(l) != 0 {
			kvFields := strings.Fields(l)
			for _, f := range kvFields {
				pField := strings.Split(f, ":")
				if len(pField) == 2 {
					setField(&block, pField[0], pField[1])
				} else {
					return output, errors.New(fmt.Sprintf("Wrong number of fields in: %v", pField))
				}
			}
			continue
		}

		if len(l) == 0 {
			output = append(output, block)
			block = Passport{}
			continue
		}
	}
	if (Passport{}) != block {
		output = append(output, block)
	}
	return output, scanner.Err()
}

func (s *Solution) part1(input []Passport) string {
	valid := 0

	for _, p := range input {
		isValid := true
		fields := reflect.TypeOf(p)
		values := reflect.ValueOf(p)
		for i := 0; i < fields.NumField(); i++ {
			f := fields.Field(i)
			fieldName, err := findFieldName(f.Tag)
			if err != nil {
				log.Printf("Unable to find field name for tag %v\n", f.Tag)
			}
			v := values.Field(i)
			if v.IsZero() && fieldName != "cid" {
				isValid = false
				break
			}
		}
		if isValid {
			valid++
		}
	}

	return fmt.Sprint(valid)
}

func (s *Solution) part2(input []Passport) string {
	valid := 0

	hairColourRegex := regexp.MustCompile(`^#[0-9a-z]{6}$`)
	eyeColourRegex := regexp.MustCompile(`^amb|blu|brn|gry|grn|hzl|oth$`)
	passportIDRegex := regexp.MustCompile(`^[0-9]{9}$`)
	heightRegex := regexp.MustCompile(`^([1-9][0-9]*)(cm|in)$`)

	for _, p := range input {
		if p.BirthYear < 1920 || p.BirthYear > 2002 {
			continue
		}

		if p.IssueYear < 2010 || p.IssueYear > 2020 {
			continue
		}

		if p.ExpirationYear < 2020 || p.ExpirationYear > 2030 {
			continue
		}

		if !hairColourRegex.MatchString(p.HairColour) {
			continue
		}

		if !eyeColourRegex.MatchString(p.EyeColour) {
			continue
		}

		if !passportIDRegex.MatchString(p.PassportID) {
			continue
		}

		heightFields := heightRegex.FindStringSubmatch(p.Height)
		if len(heightFields) != 3 {
			continue
		}

		heightNum, err := strconv.Atoi(heightFields[1])
		if err != nil {
			log.Printf("Unable to convert %s into int", heightFields[1])
		}

		if heightFields[2] == "in" && ( heightNum < 59 || heightNum > 76) {
			continue
		}

		if heightFields[2] == "cm" && ( heightNum < 150 || heightNum > 193) {
			continue
		}

		valid++

	}

	return fmt.Sprint(valid)
}

func setField(item interface{}, fieldName string, value string) error {
	v := reflect.ValueOf(item).Elem()
	if !v.CanAddr() {
		return fmt.Errorf("cannot assign to the item passed, item must be a pointer in order to assign")
	}

	fieldNames := map[string]int{}
	for i := 0; i < v.NumField(); i++ {
		typeField := v.Type().Field(i)
		tag := typeField.Tag
		pName, _ := findFieldName(tag)
		fieldNames[pName] = i
	}

	fieldNum, ok := fieldNames[fieldName]
	if !ok {
		return fmt.Errorf("field %s does not exist within the provided item", fieldName)
	}
	fieldVal := v.Field(fieldNum)
	switch fieldVal.Kind() {
	case reflect.Int:
		x, err := strconv.Atoi(value)
		if err != nil {
			return fmt.Errorf("Unable to convert %s into integer", value)
		}
		kValue := reflect.ValueOf(x)
		fieldVal.Set(kValue)
	case reflect.Ptr:
		if fieldVal.Type().Elem().Kind() == reflect.Int {
			x, err := strconv.Atoi(value)
			if err != nil {
				return fmt.Errorf("Unable to convert %s into integer", value)
			}
			kValue := reflect.ValueOf(&x)
			fieldVal.Set(kValue)
		}
	case reflect.String:
		kValue := reflect.ValueOf(value)
		fieldVal.Set(kValue)
	default:
		fmt.Printf("unexpected type %T\n", value)
	}

	return nil
}

func findFieldName(t reflect.StructTag) (string, error) {
	if name, ok := t.Lookup("passport"); ok {
		return name, nil
	}
	return "", fmt.Errorf("tag provided does not define a passport tag: %s", t)
}

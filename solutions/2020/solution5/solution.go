package solution5

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"

	"github.com/pkg/errors"
	"github.com/sazap10/advent-of-code/pkg/solution"
)

// Solution holds any data for the program
type Solution struct {
	solution.Label
}

// SeatLocation holds data for a seat location
type SeatLocation struct {
	Row    int
	Column int
	ID     int
}

// RowColumnToID gets the ID of a SeatLocation from its row and column
func (l SeatLocation) RowColumnToID() int {
	return (l.Row * 8) + l.Column
}

// ByID implements sort.Interface for []SeatLocation based on
// the ID function.
type ByID []SeatLocation

func (a ByID) Len() int           { return len(a) }
func (a ByID) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByID) Less(i, j int) bool { return a[i].ID < a[j].ID }

// New instantiates the solution
func New() *Solution {
	label := solution.NewLabel("Binary Boarding", "https://adventofcode.com/2020/day/5", "2020")
	return &Solution{
		Label: label,
	}
}

// Run contains the logic to solving the problem
func (s *Solution) Run() (string, error) {
	sampleInput, err := s.getInput("solutions/2020/solution5/sample.txt")

	if err != nil {
		return "", errors.Wrapf(err, "Unable to read sample input file")
	}

	input, err := s.getInput("solutions/2020/solution5/input.txt")

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

func (s *Solution) getInput(path string) ([]SeatLocation, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []SeatLocation
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		boardingPass := scanner.Text()
		seatLocation, err := boardingPassToSeatLocation(boardingPass)
		if err != nil {
			log.Printf("Unable to convert boarding pass to seat location %s\n", boardingPass)
		}
		lines = append(lines, seatLocation)
	}
	sort.Sort(ByID(lines))
	return lines, scanner.Err()
}

func (s *Solution) part1(input []SeatLocation) string {
	maxID := input[len(input)-1].ID

	return fmt.Sprint(maxID)
}

func (s *Solution) part2(input []SeatLocation) string {
	seatID := findMissing(input, len(input))

	return fmt.Sprint(seatID)
}

func boardingPassToSeatLocation(boardingPass string) (SeatLocation, error) {
	if len(boardingPass) != 10 {
		return SeatLocation{}, errors.New("Boarding pass string is too long, should be 10 characters")
	}

	seatLocation := SeatLocation{
		Row:    findRow(boardingPass[:7]),
		Column: findColumn(boardingPass[7:]),
	}

	seatLocation.ID = seatLocation.RowColumnToID()

	return seatLocation, nil
}

func findRow(boardingPassRow string) int {
	currentRow := struct {
		MinRow float64
		MaxRow float64
	}{
		MinRow: 0,
		MaxRow: 127,
	}
	row := 0

	for i, r := range boardingPassRow {
		switch r {
		case 'F':
			if i == (len(boardingPassRow) - 1) {
				row = int(currentRow.MinRow)
			}
			currentRow.MaxRow = math.Floor((currentRow.MinRow + currentRow.MaxRow) / 2)
		case 'B':
			if i == (len(boardingPassRow) - 1) {
				row = int(currentRow.MaxRow)
			}
			currentRow.MinRow = math.Ceil((currentRow.MinRow + currentRow.MaxRow) / 2)
		}
	}

	return row
}

func findColumn(boardingPassColumn string) int {
	currentcolumn := struct {
		Min float64
		Max float64
	}{
		Min: 0,
		Max: 7,
	}
	column := 0

	for i, r := range boardingPassColumn {
		switch r {
		case 'L':
			if i == (len(boardingPassColumn) - 1) {
				column = int(currentcolumn.Min)
			}
			currentcolumn.Max = math.Floor((currentcolumn.Min + currentcolumn.Max) / 2)
		case 'R':
			if i == (len(boardingPassColumn) - 1) {
				column = int(currentcolumn.Max)
			}
			currentcolumn.Min = math.Ceil((currentcolumn.Min + currentcolumn.Max) / 2)
		}
	}

	return column
}

func findMissing(arr []SeatLocation, size int) int {
	a := 0
	b := size - 1
	mid := 0
	for b > 1 {
		mid = a + (b-a)/2
		if arr[mid].ID-mid == arr[0].ID {
			if arr[mid+1].ID-arr[mid].ID > 1 {
				return arr[mid].ID + 1
			}
			a = mid + 1
		} else {
			if arr[mid].ID-arr[mid-1].ID > 1 {
				return arr[mid].ID - 1
			}
			b = mid - 1
		}
	}
	return -1
}

package solution_1

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sazap10/advent-of-code/pkg/solution"
)

type Solution struct {
	solution.Label
}

func New() *Solution {
	label := solution.NewLabel("Multiply numbers which sum to 2020", "https://adventofcode.com/2020/day/1", "2020")
	return &Solution{
		Label: label,
	}
}

func (s *Solution) Run() (error, string) {

	return nil, "tbd"
}

func (s *Solution) getInput() (error, []int) {
	inputURL := fmt.Sprintf("%s/input", s.Url)
	resp, err := http.Get(inputURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	log.Println("Response status:", resp.Status)
	return nil, []int{1}
}

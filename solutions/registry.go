package solutions

import (
	"github.com/sazap10/advent-of-code/pkg/solution"
	"github.com/sazap10/advent-of-code/solutions/2020/solution1"
	"github.com/sazap10/advent-of-code/solutions/2020/solution2"
	"github.com/sazap10/advent-of-code/solutions/2020/solution3"
	"github.com/sazap10/advent-of-code/solutions/2020/solution4"
)

// NewSolutions2020 generates the list of solutions to run for 2020 problems
func NewSolutions2020() solution.Map {
	output := solution.Map{}

	addSolution(output, solution1.New())
	addSolution(output, solution2.New())
	addSolution(output, solution3.New())
	addSolution(output, solution4.New())

	return output
}

func addSolution(list solution.Map, newSolution solution.Solution) {
	list[newSolution.Name()] = newSolution
}

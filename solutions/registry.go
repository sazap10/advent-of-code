package solutions

import (
	"github.com/sazap10/advent-of-code/pkg/solution"
	"github.com/sazap10/advent-of-code/solutions/2020/solution_1"
)

func NewSolutions2020() solution.Map {
	output := solution.Map{}

	addSolution(output, solution_1.New())

	return output
}

func addSolution(list solution.Map, newSolution solution.Solution) {
	list[newSolution.Name()] = newSolution
}

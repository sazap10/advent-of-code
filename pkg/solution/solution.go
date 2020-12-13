package solution

type Map map[string]Solution

type Solution interface {
	Name() string
	URL() string
	Year() string
	Run() (string, error)
}

type Label struct {
	SolutionName string
	ProblemURL   string
	SolutionYear string
}

func NewLabel(name, url, year string) Label {
	return Label{
		SolutionName: name,
		ProblemURL:   url,
		SolutionYear: year,
	}
}

// Name provides the name of the solution
func (l *Label) Name() string {
	return l.SolutionName
}

// URL provides the url to the problem
func (l *Label) URL() string {
	return l.ProblemURL
}

// Year provides the year the solution is for
func (l *Label) Year() string {
	return l.SolutionYear
}

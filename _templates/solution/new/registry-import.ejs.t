---
to: solutions/registry.go
inject: true
skip_if: github.com/sazap10/advent-of-code/solutions/<%= year %>/solution<%= day %>
after: "github.com/sazap10/advent-of-code/solutions/<%= year %>/solution<%= day - 1 %>"
eof_last: false
---
  "github.com/sazap10/advent-of-code/solutions/<%= year %>/solution<%= day %>"

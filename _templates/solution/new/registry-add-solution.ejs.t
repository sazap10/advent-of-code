---
to: solutions/registry.go
inject: true
skip_if: addSolution(output, solution<%= day %>.New())
after: addSolution\(output, solution<%= day - 1 %>\.New\(\)\)
eof_last: false
---
  addSolution(output, solution<%= day %>.New())

// see types of prompts:
// https://github.com/enquirer/enquirer/tree/master/examples
//
module.exports = [
  {
    type: 'input',
    name: 'year',
    message: "Which year is your solution for?"
  },
  {
    type: 'input',
    name: 'day',
    message: "Which day is your solution for?"
  },
  {
    type: 'input',
    name: 'name',
    message: "What is the name of the problem?"
  }
]

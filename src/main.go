package main

import "fmt"

type sValue struct {
	coldID     string
	groupType  bool
	conditions map[string]string
}

func (s sValue) String() string {
	return fmt.Sprintf("colID: %s", s.coldID)
}

func main() {

	var cold1 sValue
	cold1.coldID = "col1"
	cold1.groupType = true
	cold1.conditions = map[string]string{
		"greatherThan": "2020-08-01",
		"lessThan":     "2021-08-01",
	}

	stateValues := []sValue{
		cold1,
	}

	fmt.Println(stateValues)

	fmt.Println("hello world")
}

// {code: }
// right noe key value
/**

input


// state value
[
	{
		colId: col1,
		grouptype: true,
		conditions: {
			greatherThan: value
			lessThan: value
		}
	},
	colIdl col2,
	groupType: false,
	conditions: {
		equals:
	}
]

// have another function that becomes state to filterValues



field and conditions on its
- onChange(value, colId, type)
-

- Result
{
	colId: {
		condition1: {
			value: "2020-08-06",
			type: "greatherThan",
			filterType: "date"
		},
		// if I do not modified value2
		condition2: {
			value: "2021-08-06",
			type: "lessThan",
			filterType: "date"
		},
	}
}

it should become on later
[
	{
		field: colId,
		value: '2021-08-02',
		op: 'lessThan,
	},
	{
		field: colId,
		value: '2020-08-02',
		op: 'greatherThan,
	}
]
What I will need
- something that tells me the condition value and its value
- something that tells me the condition type

*/

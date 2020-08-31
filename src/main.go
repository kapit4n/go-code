package main

import "fmt"

type fValue struct {
	field string
	op    string
	value string
}

type sValue struct {
	colID      string
	groupType  bool
	conditions map[string]string
}

func (s sValue) String() string {
	return fmt.Sprintf("colID: %s", s.colID)
}

func main() {

	var col1 sValue
	col1.colID = "col1"
	col1.groupType = true
	col1.conditions = map[string]string{
		"greatherThan": "2020-08-01",
		"lessThan":     "2021-08-01",
	}

	var col2 sValue
	col2.colID = "col2"
	col2.groupType = false
	col2.conditions = map[string]string{
		"equals": "value",
	}

	stateValues := []sValue{
		col1, col2,
	}

	fValues := []fValue{}

	// convert stateValues to server values
	for _, value := range stateValues {
		for k := range value.conditions {
			var fVal fValue
			fVal.field = value.colID
			fVal.op = k
			fVal.value = value.conditions[k]
			fValues = append(fValues, fVal)
		}
	}

	for _, value := range fValues {
		fmt.Println(value)
	}

	fmt.Println("---------------------------")
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

// Code generated by "stringer -type=Gender -output=constants_generated.go"; DO NOT EDIT.

package users

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Male-0]
	_ = x[Female-1]
}

const _Gender_name = "MaleFemale"

var _Gender_index = [...]uint8{0, 4, 10}

func (i Gender) String() string {
	if i >= Gender(len(_Gender_index)-1) {
		return "Gender(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Gender_name[_Gender_index[i]:_Gender_index[i+1]]
}

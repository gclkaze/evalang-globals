package parameters

import "github.com/gclkaze/evalang-globals/globals"

// the variable parameter value
type VariableParameterValue struct {
	Name     string
	value    *globals.ParameterValue
	isHidden bool
}

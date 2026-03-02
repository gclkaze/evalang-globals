package parameters

import (
	"github.com/gclkaze/evalang-globals/globals"
)

type IVariableParameterValue interface {
	IsReference() bool
	Copy() globals.ParameterValue
	IsHidden() bool
	SetIsHidden(isHidden bool)
	//	OwnsReference(s stackvalue.StackValue) bool
	IsEmpty() bool
	GetName() string
	GetValue() any
	GetStringValue() string
	Equals(p *globals.ParameterValue) bool
	GetType() globals.StatementParameterTypeBase
	GetValueType() globals.StatementParameterTypeBase
	SetValue(value globals.ParameterValue)
	Length() int
	//	Assign(value stackvalue.StackValue) error
	SetPrimitiveValue(value any, t globals.StatementParameterTypeBase)
	Dump()
	Clear()
}

package globals

type ParameterValue interface {
	ISecret
	Copy() ParameterValue
	GetValue() any
	GetType() StatementParameterTypeBase
	Dump()
	Equals(*ParameterValue) bool
	GetStringValue() string
	Clear()
	IsEmpty() bool
	Length() int
	IsReference() bool
}

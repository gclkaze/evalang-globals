package stackvalue

import (
	"fmt"

	"github.com/gclkaze/evalang-globals/globals/tafargumentlistenererrortypes"
)

type StackValueType int

const (
	STRING StackValueType = iota
	INTEGER
	DOUBLE
	BOOL
	NULL
	JSON_OBJECT
	JSON_ARRAY
	BROWSER
	DB
	REST
	REFERENCE
	USER_DEFINED
)

func (e StackValueType) String() string {
	switch e {
	/*	case BROWSER:
			return "Browser"
		case DB:
			return "DB"*/
	case JSON_ARRAY:
		return "JSONArray"
	case JSON_OBJECT:
		return "JSONObject"
	case STRING:
		return "String"
	case INTEGER:
		return "Integer"
	case DOUBLE:
		return "Double"
	case BOOL:
		return "Bool"
	case NULL:
		return "Null"
	case BROWSER:
		return "Browser"
	case DB:
		return "DB"
	case REST:
		return "REST"
	case REFERENCE:
		return "Reference"
	default:
		return fmt.Sprintf("%d", int(e))
	}
}

type StackValue interface {
	ToInteger() (result int64, err error)
	ToDouble() (result float64, err error)
	ToBoolean() (result bool, err error)

	GetType() StackValueType
	Copy() StackValue
	IsTruthy() bool
	Equals(other StackValue) bool
	//ToJson() *JSONStackValue
	IsScalar() bool
	Length() (result int, err error)
	ToString() string

	Add(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error)
	Sub(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error)
	Mul(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error)
	Div(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error)
	Mod(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error)

	LesserThan(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error)
	LesserThanEqual(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error)
	Equal(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error)
	Unequal(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error)
	GreaterThan(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error)
	GreaterThanEqual(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error)

	And(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error)
	Or(s StackValue) (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error)
	Not() (result StackValue, errorType tafargumentlistenererrortypes.TAFArgumentListenerError, err error)

	ToUintPtr() uintptr
}

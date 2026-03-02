package utils

import (
	"math"
	"strconv"

	"github.com/gclkaze/evalang-globals/globals"
)

func IsString(v interface{}) (res bool, value string) {
	if v == nil {
		return false, ""
	}
	switch v.(type) {
	case string:
		return true, v.(string)
	default:
		return false, ""
	}
}

func IsDouble(v interface{}) (res bool, value float64) {
	if v == nil {
		return false, -1
	}
	_dVal, err := v.(string)
	if !err {
		dVal, err := strconv.ParseFloat(_dVal, 64)
		if err != nil {
			return false, -1
		}
		return true, dVal
	}
	return false, -1
}

func IsJson(v interface{}) (res bool, value globals.JSONStruct) {
	if v == nil {
		return false, nil
	}
	_v, err := v.(globals.JSONStruct)
	if !err {
		return false, nil
	}
	return true, _v
}

func IsNumber(v interface{}) (res bool, value float64, t globals.StatementParameterTypeBase) {

	if v == nil {
		return false, -1, globals.NULL
	}
	switch v.(type) {
	case float64:

		f := v.(float64)
		if math.Mod(f, 1.0) == 0 {
			return true, f, globals.INTEGER
		}

		return true, f, globals.DOUBLE
	case int64:
		f := v.(int64)

		return true, float64(f), globals.INTEGER

	}

	return false, -1, globals.NULL
}

func IsBool(v interface{}) (res bool, value bool) {
	if v == nil {
		return false, false
	}
	_bVal, err := v.(string)
	bVal := false
	if !err {
		bVal, err := strconv.ParseBool(_bVal)
		if err != nil {
			return false, false
		}
		return bVal, true
	}
	return bVal, false
}

func IsJSONArray(v interface{}) (res bool, t globals.JSONArrayType) {
	switch v.(type) {
	case globals.JSONGenericArray:
		{
			vv := v.(globals.JSONGenericArray)
			if len(vv) == 0 {
				return true, globals.JSON_EMPTY_ARRAY
			}
			value := vv[0]
			isString, _ := IsString(value)
			if isString {
				return true, globals.JSON_STRING_ARRAY
			}

			isBool, _ := IsBool(value)
			if isBool {
				return true, globals.JSON_BOOL_ARRAY
			}

			isNumber, _, t := IsNumber(value)
			if isNumber {
				if t == globals.INTEGER {
					return true, globals.JSON_INT_ARRAY
				}
				return true, globals.JSON_DOUBLE_ARRAY
			}

			isJSON, _ := IsJson(value)
			if isJSON {
				return true, globals.JSON_OBJECT_ARRAY
			}
		}
	default:

		return false, globals.JSON_EMPTY_ARRAY
	}
	panic(12)
	return false, globals.JSON_EMPTY_ARRAY
}

package utils

import (
	"fmt"
	"regexp"
	"strconv"

	log "github.com/sirupsen/logrus"

	jsoniter "github.com/json-iterator/go"
	"github.com/nsf/jsondiff"

	"github.com/gclkaze/evalang-globals/globals"
)

type JSONUtils struct {
	VariableExpressionRegexMatcher *regexp.Regexp
	VariableExpressionRegexGrouper *regexp.Regexp
	IndexGrouper                   *regexp.Regexp
	IndexVarGrouper                *regexp.Regexp
	EvalVarExpressionRegexGrouper  *regexp.Regexp
	Verbose                        bool
	Dbg                            bool
}

func NewJSONUtils() *JSONUtils {
	inst := &JSONUtils{}
	inst.init()
	return inst
}

func IsIntegral(val float64) bool {
	return val == float64(int(val))
}

func (s *JSONUtils) init() {
	variableExpressionMatchRegex := `^\s*[\\$][a-zA-Z_][a-zA-Z_0-9]*[\s]*([\\[][\s]*([1-9][0-9]*|[0])[\s]*[\]])*(?:\s*\.\s*[a-zA-Z_][a-zA-Z_0-9]*\s*([\\[][\s]*([1-9][0-9]*|[0])[\s]*[\]])*)*$`
	s.VariableExpressionRegexMatcher = regexp.MustCompile(variableExpressionMatchRegex)
	s.VariableExpressionRegexGrouper = regexp.MustCompile(`(?:(?P<VAR>([\\$][a-z_A-Z][a-zA-Z0-9]*))(?P<FIRST_INDEX>(\[([\s]*[1-9][0-9]*|[0])[\s]*\]))?|(?P<INDEX>[a-z_A-Z][a-zA-Z0-9]*\s*\[([\s]*[1-9][0-9]*|[0])[\s]*\])|(?P<PROPERTY>[a-z_A-Z][a-zA-Z0-9]*))+?`)
	s.IndexGrouper = regexp.MustCompile(`^\s*\w*\[[\s]*([1-9][0-9]*|[0])[\s]*\]\s*`)
	s.IndexVarGrouper = regexp.MustCompile(`^\s*(\w*)\s*\s*`)

	s.EvalVarExpressionRegexGrouper = regexp.MustCompile(`(?P<VAROCCUR>[\(]\s*(?P<VAR>[$][\w][\w\.\[\]\s]+)\s*[)])`)

}

func (s *JSONUtils) GetJSONVariableExpressionTokens(a *string) (res bool, myVar *[]globals.PathElem) {

	var tokens []string

	var p []globals.PathElem
	groupNames := s.VariableExpressionRegexGrouper.SubexpNames()
	for matchNum, match := range s.VariableExpressionRegexGrouper.FindAllStringSubmatch(*a, -1) {
		for groupIdx, group := range match {
			name := groupNames[groupIdx]
			if name == "" {
				name = "*"
				continue
			}
			if group == "" {
				continue
			}
			if s.Verbose {
				log.Printf("#%d text: '%s', group: '%s'\n", matchNum, group, name)
			}
			tokens = append(tokens, group)

			var v globals.PathElem
			switch {
			case name == "VAR":
				{
					v = globals.PathElem{Value: group, Index: -1, Type: globals.VARIABLE_DECL}
					break
				}
			case name == "PROPERTY":
				{
					v = globals.PathElem{Value: group, Index: -1, Type: globals.PROPERTY}
					break
				}
			case name == "INDEX":
				{
					res, index := s.GetIndex(&group)
					if !res {
						return false, nil
					}
					res, varName := s.GetIndexVar(&group)
					if !res {
						return false, nil
					}
					v = globals.PathElem{Value: varName, Index: index, Type: globals.INDEX}
					break
				}
			case name == "FIRST_INDEX":
				{
					res, index := s.GetIndex(&group)
					if !res {
						return false, nil
					}
					res, varName := s.GetIndexVar(&group)
					if !res {
						return false, nil
					}
					v = globals.PathElem{Value: varName, Index: index, Type: globals.INDEX}
					break
				}

			default:
				{
					log.Printf("Unknown token captured %s", name)
					continue
				}
			}
			p = append(p, v)
		}
	}

	if s.Dbg {
		fmt.Printf("tokens: %v\n", tokens)
		fmt.Printf("p: %v\n", p)
	}
	//return true, &parameters.JsonParameterValue{CurrentPath: tokens, Name: tokens[0], Path: &p}
	return true, &p
}

func (s *JSONUtils) GetIndex(a *string) (res bool, value int) {
	ss := *a
	v := s.IndexGrouper.FindAllStringSubmatch(ss, -1)
	if len(v) == 0 || len(v) == 1 && len(v[0]) != 2 {
		return false, -1
	}
	i, err := strconv.ParseInt(v[0][1], 10, 64)
	if err != nil {
		if s.Dbg {
			log.Println("Couldn't parse INDEX " + ss)
		}
		return false, -1
	}
	return true, int(i)
}

func (s *JSONUtils) GetIndexVar(a *string) (res bool, value string) {
	ss := *a
	v := s.IndexVarGrouper.FindAllStringSubmatch(ss, -1)
	if len(v) == 0 || len(v) == 1 && len(v[0]) != 2 {
		return false, ""
	}
	return true, v[0][1]
}

/*
func JSONDiff(first globals.JSONStruct, second globals.JSONStruct) bool {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	currentString, err := json.Marshal(first)
	if err != nil {
		log.Errorf(err.Error())
		return false
	}
	otherString, err := json.Marshal(second)
	if err != nil {
		log.Errorf(err.Error())
		return false
	}
	opt := jsondiff.DefaultJSONOptions()

	diff, diffString := jsondiff.Compare([]byte(currentString), []byte(otherString), &opt)
	if diff == jsondiff.FullMatch {
		return true
	}
	log.Debug(diffString)
	return false
}*/

func JSONDiff(first globals.JSONObjectGen, second globals.JSONObjectGen) bool {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	currentString, err := json.Marshal(first)
	if err != nil {
		log.Errorf("%s", err.Error())
		return false
	}
	otherString, err := json.Marshal(second)
	if err != nil {
		log.Errorf("%s", err.Error())
		return false
	}
	opt := jsondiff.DefaultJSONOptions()

	diff, diffString := jsondiff.Compare([]byte(currentString), []byte(otherString), &opt)
	if diff == jsondiff.FullMatch {
		return true
	}
	log.Println(diffString)
	return false //s.value == strV.value
}

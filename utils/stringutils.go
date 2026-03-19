package utils

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"unicode"

	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"

	"github.com/gclkaze/evalang-globals/globals"
)

func JSONCopy(o globals.JSONObjectGen) (globals.JSONObjectGen, error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	b, err := json.Marshal(o)
	if err != nil {
		return nil, err
	}

	var dst interface{}
	err = json.Unmarshal(b, &dst)
	if err != nil {
		return nil, err
	}
	return dst, nil
}

func JSONArrayCopy(o globals.JSONArrayGen) (globals.JSONArrayGen, error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	b, err := json.Marshal(o)
	if err != nil {
		return nil, err
	}

	var dst []interface{}
	err = json.Unmarshal(b, &dst)
	if err != nil {
		return nil, err
	}
	return dst, nil
}

func ContainsString(content string, searchString string) (value bool, err error) {
	return strings.Contains(content, searchString), nil
}

func StarsWith(content string, searchString string) (value bool, err error) {
	return strings.HasPrefix(content, searchString), nil
}

func EndsWith(content string, searchString string) (value bool, err error) {
	return strings.HasSuffix(content, searchString), nil
}

func Trim(content string) (res string, err error) {
	return strings.TrimSpace(content), nil
}

func TrimLeft(content string) (res string, err error) {
	return strings.TrimLeft(content, " "), nil
}
func TrimRight(content string) (res string, err error) {
	return strings.TrimRight(content, " "), nil
}

func ReplaceAllString(content string, old string, replacement string) (res string) {
	return strings.ReplaceAll(content, old, replacement)
}

func GetAbsolutePath(filename string) (path string, err error) {
	return filepath.Abs(filename)
}

func GetFirstIndexAndPath(expr string) (ind int, path string, err error) {
	tokens := strings.Split(expr, ".")
	fmt.Printf("tokens: %v\n", tokens)
	indx, err := strconv.Atoi(tokens[1])
	if err != nil {
		return -1, "", err
	}
	p := ""
	for i := 2; i < len(tokens); i++ {
		if i == 2 {
			p += tokens[i]
		} else {
			p += "." + tokens[i]
		}
	}
	return indx, p, nil
}

func GetVarComponents(expr string) (v string, ind int, path string, err error) {
	e := ReplaceBracketsWithDotsInExpr(expr)
	tokens := strings.Split(e, ".")
	fmt.Printf("tokens: %v\n", tokens)
	indx, err := strconv.Atoi(tokens[1])
	if err != nil {
		return "", -1, "", err
	}
	p := ""
	for i := 2; i < len(tokens); i++ {
		if i == 2 {
			p += tokens[i]
		} else {
			p += "." + tokens[i]
		}
	}
	return tokens[0], indx, p, nil
}

func GetPathFromExpressionAndVarname(varName string, expr string) string {
	if strings.HasPrefix(expr, varName+".") {
		path := strings.Replace(expr, varName+".", "", 1)
		return path
	}
	return expr
}

func ReplaceBracketsWithDotsInExpr(s string) string {
	if s == "" {
		return s
	}
	s = strings.ReplaceAll(s, "[", ".")
	s = strings.ReplaceAll(s, "]", ".")
	s = strings.ReplaceAll(s, "..", ".")
	if strings.HasSuffix(s, ".") {
		return s[:len(s)-1]
	}
	return s
}

func GetBracketedString(s string) string {
	var sb strings.Builder
	for _, r := range s {
		sb.WriteString("[")
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			sb.WriteByte('\\')
		}
		sb.WriteRune(r)
		sb.WriteString("]")
	}
	return sb.String()
}

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	if err == nil {
		return true
	}
	log.Println(err)
	return false
}

func ReadFile(filename string) *string {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil
	}
	c := string(content)
	return &c
}

func WriteFile(dir string, filename string, content []byte) error {
	p := filepath.Join(dir, filename)

	err := os.WriteFile(p, content, 0600)
	if err != nil {
		return err
	}
	return err
}

func RemoveFile(path string) error {
	err := os.Remove(path)
	return err
}

func GetLimitedString(s string, limit int) string {
	if len(s) < limit+1 {
		return s
	}
	return s[:limit] + "..."
}

func XPATHReplaceEscapedSlashes(xpath *string) {
	if strings.Contains(*xpath, "\\\"") {
		*xpath = strings.ReplaceAll(*xpath, "\\\"", "\"")
	}
}

func CapitalizeOnlyFirst(s string) string {
	s = string(unicode.ToUpper(rune(s[0]))) + s[1:]
	return s
}

func TrimDoubleQuotedString(s *string) {
	*s = strings.TrimPrefix(*s, "\"")
	*s = strings.TrimSuffix(*s, "\"")
}

func TrimSingleQuotedString(s *string) {
	*s = strings.TrimPrefix(*s, "\\'")
	*s = strings.TrimSuffix(*s, "\\'")
}

func StringArrayJoin(a []string, separator string) string {
	if len(a) == 0 {
		return ""
	}
	return strings.Join(a, separator)
}

func ReadJSON(val string) (globals.JSONStruct, error) {
	XPATHReplaceEscapedSlashes(&val)

	var jsonMap globals.JSONStruct
	err := json.Unmarshal([]byte(val), &jsonMap)
	return jsonMap, err
}
func GetRandomString() string {
	return uuid.New().String()
}

func UniqueStrings(in []string) []string {
	seen := make(map[string]struct{})
	out := make([]string, 0, len(in))

	for _, v := range in {
		v = strings.TrimSpace(v)
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		out = append(out, v)
	}
	return out
}

func SQLResultToJSON(rows *sql.Rows) (globals.JSONArray, error) {
	columnTypes, err := rows.ColumnTypes()

	if err != nil {
		return nil, err
	}

	count := len(columnTypes)
	finalRows := globals.JSONArray{}

	for rows.Next() {

		scanArgs := make([]interface{}, count)

		for i, v := range columnTypes {

			switch v.DatabaseTypeName() {
			case "VARCHAR", "TEXT", "UUID", "TIMESTAMP":
				scanArgs[i] = new(sql.NullString)
			case "BOOL":
				scanArgs[i] = new(sql.NullBool)
			case "DOUBLE":
				scanArgs[i] = new(sql.NullFloat64)
			case "BIGINT":
				scanArgs[i] = new(sql.NullInt64)
			case "INT4":
				scanArgs[i] = new(sql.NullInt64)
			default:
				scanArgs[i] = new(sql.NullString)
			}
		}

		err := rows.Scan(scanArgs...)

		if err != nil {
			return nil, err
		}

		masterData := map[string]interface{}{}

		for i, v := range columnTypes {

			columnType := v.DatabaseTypeName()
			if columnType == "BIGINT" || columnType == "INT" {
				z, ok := (scanArgs[i]).(*sql.NullInt64)
				if !ok {
					masterData[v.Name()] = scanArgs[i]
					continue
				}
				if !ok {
					masterData[v.Name()] = scanArgs[i]
					continue
				}
				masterData[v.Name()] = float64(z.Int64)
				continue
			}
			if columnType == "DOUBLE" {
				z, ok := (scanArgs[i]).(*sql.NullFloat64)
				if !ok {
					masterData[v.Name()] = scanArgs[i]
					continue
				}
				masterData[v.Name()] = float64(z.Float64)
				continue
			}
			if columnType == "DECIMAL" {
				z, ok := (scanArgs[i]).(*sql.NullString)
				if !ok {
					masterData[v.Name()] = scanArgs[i]
					continue
				}
				dV, err := strconv.ParseFloat(z.String, 64)
				if err != nil {
					masterData[v.Name()] = scanArgs[i]
					continue
				}
				masterData[v.Name()] = float64(dV)
				continue
			}

			if columnType == "BOOL" {
				z, ok := (scanArgs[i]).(*sql.NullString)
				if !ok {
					masterData[v.Name()] = scanArgs[i]
					continue
				}
				bV, err := strconv.ParseBool(z.String)
				if err != nil {
					masterData[v.Name()] = scanArgs[i]
					continue
				}
				masterData[v.Name()] = bool(bV)
				continue
			}
			if columnType == "VARCHAR" || columnType == "TEXT" {
				z, ok := (scanArgs[i]).(*sql.NullString)
				if !ok {
					masterData[v.Name()] = scanArgs[i]
					continue
				}

				masterData[v.Name()] = z.String
				continue
			}
		}

		finalRows = append(finalRows, masterData)
	}

	/*	z, err := json.Marshal(finalRows)
		fmt.Printf("z: %v\n", z)
	*/
	return finalRows, nil
}

func FileOrString(s *string) {
	if s == nil {
		return
	}
	if _, err := os.Stat(*s); errors.Is(err, os.ErrNotExist) {
		return
	} else {
		//its a path thus we need to read it and assign to the Arg
		content := ReadFile(*s)
		if content == nil {
			return
		}
		*s = *content
	}
}

package utils

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/xwb1989/sqlparser"
)

func GetQueries(s string) ([]string, error) {
	r := strings.NewReader(s)
	tokens := sqlparser.NewTokenizer(r)
	var queries []string
	for {
		stmt, err := sqlparser.ParseNext(tokens)
		if err == io.EOF {
			break
		}
		// Do something with stmt or err.
		if err != nil {
			return []string{}, err
		}
		q := sqlparser.String(stmt)
		queries = append(queries, q)
	}

	return queries, nil
}

func GetReadQuery(s string) (string, error) {
	r := strings.NewReader(s)
	tokens := sqlparser.NewTokenizer(r)

	// Otherwise do something with stmt
	/*	switch stmt := stmt.(type) {
		case *sqlparser.Select:
			_ = stmt
			return s, nil
		}
		return "", errors.New("'" + s + "' is not a read query")
	*/

	amount := 0
	var queries []string
	for {
		stmt, err := sqlparser.ParseNext(tokens)
		if err == io.EOF {
			break
		}
		// Do something with stmt or err.
		if err != nil {
			return "", err
		}
		switch stmt := stmt.(type) {
		case *sqlparser.Select:
			q := sqlparser.String(stmt)
			queries = append(queries, q)
			_ = stmt
		}

		amount++
	}

	if amount > 1 {
		return "", fmt.Errorf("%d queries were provided instead of 1", amount)
	}
	return queries[0], nil
}

func GetWriteQueries(s string) ([]string, error) {
	r := strings.NewReader(s)
	tokens := sqlparser.NewTokenizer(r)
	var queries []string
	for {
		stmt, err := sqlparser.ParseNext(tokens)
		if err == io.EOF {
			break
		}
		// Do something with stmt or err.
		if err != nil {
			return []string{}, err
		}
		q := sqlparser.String(stmt)

		switch stmt := stmt.(type) {
		case *sqlparser.Select:
			_ = stmt
			return []string{}, errors.New("'" + q + "' is not a write query")
		}
		queries = append(queries, q)
	}

	return queries, nil
}

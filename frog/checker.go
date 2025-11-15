// Copyright (C) by abdenour souane
// you have a right to modify it upgrade it or do whatever you want
// but u have to keep my name on it
package frog

import (
	"strings"
)

type NumberType int

const (
	Integer NumberType = iota
	Float
)

func CheckNumberType(number string) NumberType {
	var splited []string = strings.SplitAfter(number, ".")
	if len(splited) == 1 {
		return Integer
	} else {
		return Float
	}
}

func ExpectInteger(number string) bool {
	return CheckNumberType(number) == Integer
}

func ExpectFloat(number string) bool {
	return CheckNumberType(number) == Float
}

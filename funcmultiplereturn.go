package main

import "strings"

func listNames(fullName string) (string, string) {

	name := strings.ToUpper(fullName)
    names := strings.Split(name, " ")
	initials := []string{}

	for _, name := range names {

		initials = append(initials, name)

	}

	if len(initials) > 1 {

		return initials[0], initials[1]
	}

	return initials[0], "_"
}
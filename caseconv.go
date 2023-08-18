// Package caseconv provides utility function to convert strings into different casings.
package caseconv

import (
	"strings"
)

// toLower returns a copy of the input string with all characters set to toLower.
func toLower(s string) string {
	return strings.ToLower(s)
}

// toUpper returns a copy of the input string with all characters set to toUpper.
func toUpper(s string) string {
	return strings.ToUpper(s)
}

// toCapitalized returns a copy of the input string with the first character set to uppercase.
func toCapitalized(s string) string {
	parts := strings.Split(s, "")
	return toUpper(parts[0]) + strings.Join(parts[1:], "")
}

// toCapitalizedWords takes a string which is assumed to be me zero or more space-separated
// groupings of characters.
//
// A new string is returned which returns these grouping of characters with the first letter of each
// grouping set to uppercase.
func toCapitalizedWords(s string) string {
	words := strings.Split(s, " ")
	transformed := make([]string, 0)

	for _, w := range words {
		res := toCapitalized(w)
		transformed = append(transformed, res)
	}

	return strings.Join(transformed, " ")
}

// toDecapitalized returns a copy of the input string with the first character of the string set to
// lowercase.
func toDecapitalized(s string) string {
	parts := strings.Split(s, "")
	return toLower(parts[0]) + strings.Join(parts[1:], "")
}

func replace(s, old, new string) string {
	return strings.ReplaceAll(s, old, new)
}

// stripDashes returns a copy of the input string with all dashes replaced by a whitespace
// character.
func stripDashes(s string) string {
	return replace(s, "-", " ")
}

// stripUnderscores returns a copy of the input string with all underscores replaced by a whitespace
// character.
func stripUnderscores(s string) string {
	return replace(s, "_", " ")
}

// stripWhitespace returns a copy of the input string with all whitespace characters removed.
func stripWhitespace(s string) string {
	return replace(s, " ", "")
}

// strip returns a copy of the input string with all dashes and underscores replaced by whitespace
// characters.
func strip(s string) string {
	return reducePipeline(stripUnderscores, stripDashes)(s)
}

// addDashes returns a copy of the input string with all whitespace characters replaced by dashes.
func addDashes(s string) string {
	return replace(s, " ", "-")
}

// addUnderscores returns a copy of the input string with all whitespace characters replaced by
// underscores.
func addUnderscores(s string) string {
	return replace(s, " ", "_")
}

// A StringPipeFn is any function which takes in a string and returns a string. These functions are
// intended to be used in a pipleine-like fashion.
type StringPipeFn = func(string) string

// pipe composes the two passed in functions, returning a new function which calls the functions in
// reverse order of argument order.
func pipe(a, b StringPipeFn) StringPipeFn {
	return func(s string) string {
		return b(a(s))
	}
}

// reducePipeline composes all passed in functions, returning a new function which calls the
// functions in reverse argument order.
func reducePipeline(fns ...StringPipeFn) StringPipeFn {
	acc := func(s string) string { return s }
	for _, fn := range fns {
		acc = pipe(acc, fn)
	}
	return acc
}

// ToStartCase returns a copy of the input string that is set to Start Case.
func ToStartCase(s string) string {
	return reducePipeline(strip, toCapitalizedWords)(s)
}

// ToPascalCase returns a copy of the input string that is set to PascalCase.
func ToPascalCase(s string) string {
	return reducePipeline(ToStartCase, stripWhitespace)(s)
}

// ToCamelCase returns a copy of the input string that is set to camelCase.
func ToCamelCase(s string) string {
	return reducePipeline(ToPascalCase, toDecapitalized)(s)
}

// ToKebabCase returns a copy of the input string that is set to kebab-case.
func ToKebabCase(s string) string {
	return reducePipeline(toLower, strip, addDashes)(s)
}

// ToSnakeCase returns a copy of the input string that is set to snake_case.
func ToSnakeCase(s string) string {
	return reducePipeline(toLower, strip, addUnderscores)(s)
}

// ToConstantCase returns a copy of the input string that is set to CONSTANT_CASE.
func ToConstantCase(s string) string {
	return reducePipeline(toUpper, strip, addUnderscores)(s)
}

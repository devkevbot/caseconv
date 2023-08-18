package main

import (
	"strings"
)

func main() {
}

func lowercase(s string) string {
	return strings.ToLower(s)
}

func uppercase(s string) string {
	return strings.ToUpper(s)
}

func capitalize(s string) string {
	parts := strings.Split(s, "")
	return uppercase(parts[0]) + strings.Join(parts[1:], "")
}

func capitalizeWords(s string) string {
	words := strings.Split(s, " ")
	transformed := make([]string, 0)

	for _, w := range words {
		res := capitalize(w)
		transformed = append(transformed, res)
	}

	return strings.Join(transformed, " ")
}

func decapitalize(s string) string {
	parts := strings.Split(s, "")
	return lowercase(parts[0]) + strings.Join(parts[1:], "")
}

func replace(s, from, to string) string {
	return strings.ReplaceAll(s, from, to)
}

func stripDashes(s string) string {
	return replace(s, "-", " ")
}

func stripUnderscores(s string) string {
	return replace(s, "_", " ")
}

func stripWhitespace(s string) string {
	return replace(s, " ", "")
}

func strip(s string) string {
	return transformPipe(stripUnderscores, stripDashes)(s)
}

func addDashes(s string) string {
	return replace(s, " ", "-")
}

func addUnderscores(s string) string {
	return replace(s, " ", "_")
}

type StringPipe = func(string) string

func pipe(a, b StringPipe) StringPipe {
	return func(s string) string {
		return b(a(s))
	}
}

func reduce[T, M any](s []T, f func(M, T) M, initValue M) M {
	acc := initValue
	for _, v := range s {
		acc = f(acc, v)
	}
	return acc
}

func transformPipe(fns ...StringPipe) StringPipe {
	return reduce(fns, pipe, func(s string) string { return s })
}

func StartCase(s string) string {
	return transformPipe(strip, capitalizeWords)(s)
}

func PascalCase(s string) string {
	return transformPipe(StartCase, stripWhitespace)(s)
}

func CamelCase(s string) string {
	return transformPipe(PascalCase, decapitalize)(s)
}

func KebabCase(s string) string {
	return transformPipe(lowercase, strip, addDashes)(s)
}

func SnakeCase(s string) string {
	return transformPipe(lowercase, strip, addUnderscores)(s)
}

func ConstantCase(s string) string {
	return transformPipe(uppercase, strip, addUnderscores)(s)
}

package main

import "testing"

func TestCapitalize(t *testing.T) {
	testCases := []struct {
		desc     string
		input    string
		expected string
	}{
		{
			desc:     "should capitalize when all lower",
			input:    "foobar",
			expected: "Foobar",
		},
		{
			desc:     "should capitalize when all upper",
			input:    "FOOBAR",
			expected: "FOOBAR",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			res := capitalize(tc.input)
			if res != tc.expected {
				t.Errorf("got %s, want %s", res, tc.expected)
			}
		})
	}
}

func TestCapitalizeWords(t *testing.T) {
	testCases := []struct {
		desc     string
		input    string
		expected string
	}{
		{
			desc:     "should capitalize when all lower",
			input:    "foo bar baz",
			expected: "Foo Bar Baz",
		},
		{
			desc:     "should capitalize when all upper",
			input:    "FOO BAR BAZ",
			expected: "FOO BAR BAZ",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			res := capitalizeWords(tc.input)
			if res != tc.expected {
				t.Errorf("got %s, want %s", res, tc.expected)
			}
		})
	}
}

func TestDecapitalize(t *testing.T) {
	testCases := []struct {
		desc     string
		input    string
		expected string
	}{
		{
			desc:     "should decapitalize when all lower",
			input:    "foobar",
			expected: "foobar",
		},
		{
			desc:     "should decapitalize when all upper",
			input:    "FOOBAR",
			expected: "fOOBAR",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			res := decapitalize(tc.input)
			if res != tc.expected {
				t.Errorf("got %s, want %s", res, tc.expected)
			}
		})
	}
}

func TestStartCase(t *testing.T) {
	testCases := []struct {
		desc     string
		input    string
		expected string
	}{
		{
			desc:     "should handle spaces",
			input:    "hello world",
			expected: "Hello World",
		},
		{
			desc:     "should handle underscores",
			input:    "hello_world",
			expected: "Hello World",
		},
		{
			desc:     "should handle dashes",
			input:    "hello-world",
			expected: "Hello World",
		},
		{
			desc:     "should handle casing",
			input:    "Hello World",
			expected: "Hello World",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			res := StartCase(tc.input)
			if res != tc.expected {
				t.Errorf("got %s, want %s", res, tc.expected)
			}
		})
	}
}

func TestPascalCase(t *testing.T) {
	testCases := []struct {
		desc     string
		input    string
		expected string
	}{
		{
			desc:     "should handle spaces",
			input:    "hello world",
			expected: "HelloWorld",
		},
		{
			desc:     "should handle underscores",
			input:    "hello_world",
			expected: "HelloWorld",
		},
		{
			desc:     "should handle dashes",
			input:    "hello-world",
			expected: "HelloWorld",
		},
		{
			desc:     "should handle casing",
			input:    "Hello World",
			expected: "HelloWorld",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			res := PascalCase(tc.input)
			if res != tc.expected {
				t.Errorf("got %s, want %s", res, tc.expected)
			}
		})
	}
}

func TestCamelCase(t *testing.T) {
	testCases := []struct {
		desc     string
		input    string
		expected string
	}{
		{
			desc:     "should handle spaces",
			input:    "hello world",
			expected: "helloWorld",
		},
		{
			desc:     "should handle underscores",
			input:    "hello_world",
			expected: "helloWorld",
		},
		{
			desc:     "should handle dashes",
			input:    "hello-world",
			expected: "helloWorld",
		},
		{
			desc:     "should handle casing",
			input:    "Hello World",
			expected: "helloWorld",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			res := CamelCase(tc.input)
			if res != tc.expected {
				t.Errorf("got %s, want %s", res, tc.expected)
			}
		})
	}
}

func TestKebabCase(t *testing.T) {
	testCases := []struct {
		desc     string
		input    string
		expected string
	}{
		{
			desc:     "should handle spaces",
			input:    "hello world",
			expected: "hello-world",
		},
		{
			desc:     "should handle underscores",
			input:    "hello_world",
			expected: "hello-world",
		},
		{
			desc:     "should handle dashes",
			input:    "hello-world",
			expected: "hello-world",
		},
		{
			desc:     "should handle casing",
			input:    "Hello World",
			expected: "hello-world",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			res := KebabCase(tc.input)
			if res != tc.expected {
				t.Errorf("got %s, want %s", res, tc.expected)
			}
		})
	}
}

func TestSnakeCase(t *testing.T) {
	testCases := []struct {
		desc     string
		input    string
		expected string
	}{
		{
			desc:     "should handle spaces",
			input:    "hello world",
			expected: "hello_world",
		},
		{
			desc:     "should handle underscores",
			input:    "hello_world",
			expected: "hello_world",
		},
		{
			desc:     "should handle dashes",
			input:    "hello_world",
			expected: "hello_world",
		},
		{
			desc:     "should handle casing",
			input:    "Hello World",
			expected: "hello_world",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			res := SnakeCase(tc.input)
			if res != tc.expected {
				t.Errorf("got %s, want %s", res, tc.expected)
			}
		})
	}
}

func TestConstantCase(t *testing.T) {
	testCases := []struct {
		desc     string
		input    string
		expected string
	}{
		{
			desc:     "should handle spaces",
			input:    "hello world",
			expected: "HELLO_WORLD",
		},
		{
			desc:     "should handle underscores",
			input:    "HELLO_WORLD",
			expected: "HELLO_WORLD",
		},
		{
			desc:     "should handle dashes",
			input:    "HELLO_WORLD",
			expected: "HELLO_WORLD",
		},
		{
			desc:     "should handle casing",
			input:    "Hello World",
			expected: "HELLO_WORLD",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			res := ConstantCase(tc.input)
			if res != tc.expected {
				t.Errorf("got %s, want %s", res, tc.expected)
			}
		})
	}
}

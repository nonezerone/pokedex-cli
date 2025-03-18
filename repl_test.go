package main

import "testing"

func TestCleanInput(t *testing.T) {
    cases := []struct {
        input    string
        expected []string
    }{
        {
            input:    "  hello  world  ",
            expected: []string{"hello", "world"},
        },
        {
            input:    " UPCASE WORDS HULLO",
            expected: []string{"upcase", "words", "hullo"},
        },
        {
            input:    "  RaNdOm CAsE worDs  ",
            expected: []string{"random", "case", "words"},
        },
    }

    for _, c := range cases {
        actual := cleanInput(c.input)
        expected := c.expected
        if len(actual) != len(expected) {
            t.Errorf("lengths do not match: '%v' vs '%v'", actual, expected)
            continue
        }
        for i := range actual {
            word := actual[i]
            expectedWord := c.expected[i]
            if word != expectedWord {
                t.Errorf("cleanInput(%v) == %v, expected %v", c.input, actual, c.expected)
            }
        }
    }
}

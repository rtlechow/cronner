// Tideland Go Library - String Extensions - Unit Tests
//
// Copyright (C) 2015-2016 Frank Mueller / Tideland / Oldenburg / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

package stringex_test

//--------------------
// IMPORTS
//--------------------

import (
	"strings"
	"testing"

	"github.com/tideland/golib/audit"
	"github.com/tideland/golib/stringex"
)

//--------------------
// TESTS
//--------------------

// TestSplitFilter tests splitting with a filter.
func TestSplitFilter(t *testing.T) {
	assert := audit.NewTestingAssertion(t, true)
	tests := []struct {
		name   string
		in     string
		sep    string
		filter func(string) bool
		out    []string
	}{
		{
			"all fine",
			"a/b/c",
			"/",
			func(p string) bool {
				return p != ""
			},
			[]string{"a", "b", "c"},
		}, {
			"filter empty parts",
			"/a/b//c/",
			"/",
			func(p string) bool {
				return p != ""
			},
			[]string{"a", "b", "c"},
		}, {
			"filter all parts",
			"a/b/c",
			"/",
			func(p string) bool {
				return p == "x"
			},
			[]string{},
		}, {
			"filter empty input",
			"",
			"/",
			func(p string) bool {
				return true
			},
			[]string{""},
		}, {
			"filter not splitted",
			"/a/b/c/",
			"+",
			func(p string) bool {
				return p != ""
			},
			[]string{"/a/b/c/"},
		},
	}
	for i, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Logf("splitfilter %d: %s", i, test.name)
			out := stringex.SplitFilter(test.in, test.sep, test.filter)
			assert.Equal(out, test.out)
		})
	}
}

// TestSplitMap tests spliting with a mapper.
func TestSplitMap(t *testing.T) {
	assert := audit.NewTestingAssertion(t, true)
	tests := []struct {
		name   string
		in     string
		sep    string
		mapper func(string) (string, bool)
		out    []string
	}{
		{
			"uppercase all",
			"a/b/c",
			"/",
			func(p string) (string, bool) {
				return strings.ToUpper(p), true
			},
			[]string{"A", "B", "C"},
		}, {
			"filter empty parts, uppercase the rest",
			"/a/b//c/",
			"/",
			func(p string) (string, bool) {
				if p != "" {
					return strings.ToUpper(p), true
				}
				return "", false
			},
			[]string{"A", "B", "C"},
		}, {
			"filter all parts",
			"a/b/c",
			"/",
			func(p string) (string, bool) {
				return p, p == "x"
			},
			[]string{},
		}, {
			"encapsulate even empty input",
			"",
			"/",
			func(p string) (string, bool) {
				return "(" + p + ")", true
			},
			[]string{"()"},
		}, {
			"uppercase not splitted",
			"/a/b/c/",
			"+",
			func(p string) (string, bool) {
				return strings.ToUpper(p), true
			},
			[]string{"/A/B/C/"},
		},
	}
	for i, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Logf("splitmap %d: %s", i, test.name)
			out := stringex.SplitMap(test.in, test.sep, test.mapper)
			assert.Equal(out, test.out)
		})
	}
}

// EOF

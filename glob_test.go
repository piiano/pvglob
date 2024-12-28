package pvglob

import (
	"testing"
)

func TestGlobMatch(t *testing.T) {
	for _, tc := range []struct {
		pattern string
		value   string
		match   bool
	}{
		{`hello`, `hello`, true},
		{`he*o`, `hello`, true},
		{`he?lo`, `hello`, true},
		{`he\*lo`, `he*lo`, true},
		{`he\*lo`, `hello`, false},
		{`he\?lo`, `he?lo`, true},
		{`he\?lo`, `hello`, false},
		{`*`, `anything`, true},
		{`*`, `*`, true},
		{`*`, ``, true},
		{`h?*o`, `hello`, true},
		{`h\?*o`, `h?llo`, true},
		{`he*`, `hello`, true},
		{`he*`, `zhe`, false},
		{`*o`, `hello`, true},
		{`*o`, `hell`, false},
		{`he\*lo`, `he*lo`, true},
		{"g*d*e", "goodbye", true},
		{"g*d*e", "goodbyd", false},
		{"g*d*d", "goodbye", false},
		{"*g*d*e", "goodbydgoodbye", true},
		{"g*d*e", "doodbydgoodbye", false},
		{`Melind*ue*`, `Melindique Ray`, true},
		{`Aaron *MD*`, `Aaron Austin MD `, true},
		{`Aaron *MD*`, `Aaron Matthews`, false},
		{`john?*gmail.com`, `john2@gmail.com`, true},
		{`john?*`, `john2@gmail.com`, true},
		{`john*?gmail.com`, `john2@gmail.com`, true},
		{`john*?`, `john2@gmail.com`, true},
		{`a??b`, `aXXb`, true},
		{`a??b`, `aXb`, false},
		{`a??b`, `aXXXb`, false},
		{`a?*`, `a`, false},
		{`a?*`, `ab`, true},
		{`a?*`, `abc`, true},
		{`a*?`, `a`, true},
		{`a*?`, `ab`, true},
		{`a*?`, `abc`, true},
	} {
		t.Run(tc.pattern, func(t *testing.T) {
			match := Compile(tc.pattern).Match(tc.value)
			if match != tc.match {
				t.Errorf("expected %v, got %v", tc.match, match)
			}
		})
	}
}

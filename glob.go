package pvglob

// Compile compiles a pattern to a ready-to-use object that can be used to match strings.
func Compile(pattern string) Parsed {
	return parse(pattern)
}

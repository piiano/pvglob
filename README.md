# pvglob - General-purpose glob matching for Go

A glob matching library with no special treatment for file separators.

## Install

```shell
go get -u github.com/piiano/pvglob
```

## Wildcards

pvglob supports the following wildcard characters:

* `*` - Any string in any length, including empty.
* `?` - A single character.

The backslash (`\`) is used as the escape character in search patterns. To search for the literal characters `*` and `?`, prefix them with a backslash, like this: `\*` and `\?`. You can also escape the backslash itself by using two backslashes `\\`.

## Usage example

Here's an example of how to use pvglob for pattern matching:

```go
package main

import (
    "fmt"
    "github.com/piiano/pvglob"
)

func main() {
    pattern := "foo*bar"

    // Compile the pattern once.
    matcher := pvglob.Compile(pattern)

    // Evaluate a string. In this case: wildcard matches a single character.
    matched := matcher.Match("fooXbar") // true
    
    // Wildcard matches multiple characters as well, no need for "**".
    matched = matcher.Match("foo123bar") // true

    // Wildcard evaluates to zero-characters as well.
    matched = matcher.Match("foobar") // true

    // And in this case, the evaluation fails.
    matched = matcher.Match("fizzbuzz") // false
}
```

## Matching over encrypted data

This library supports the [Piiano Vault's substring search over encrypted data](https://docs.piiano.com/guides/write-and-read-personal-data/search-objects/substring-search-objects), and provides functionality to support Trigram extraction from a search pattern, as well as additional information over each non-wildcar literal that supports correct evaluation.
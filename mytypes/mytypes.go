package mytypes

import (
	"strings"
)

type MyInt int

func (i MyInt) Twice() MyInt {
	return 2 * i
}

type MyString string

func (s MyString) Len() int {
	return len(s)
}

type MultiString []string

func (ms MultiString) Join() string {
	return strings.Join(ms, "")
}

type MyBuilder struct {
	strings.Builder
}

type StringUppercaser struct {
	strings.Builder
}

func (su StringUppercaser) Uppercase() string {
	return strings.ToUpper(su.String())
}

type IntBuilder struct {
	Contents []int
}

func (ib IntBuilder) Sum() int {
	result := 0
	for _, v := range ib.Contents {
		result += v
	}
	return result

}

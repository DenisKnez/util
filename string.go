package util

import "fmt"

type String struct{}

func (s *String) ToQuotedString(toQuote string) string {
	return fmt.Sprintf(`"%s"`, toQuote)
}

func (s *String) ToPointedString(toPoint string) *string {
	return &toPoint
}

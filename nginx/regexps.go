package nginx

import "regexp"

var (
	rMatchMultipleSpace = regexp.MustCompile("\\s\\s+")
	rMatchComments      = regexp.MustCompile("\\#.*")
)

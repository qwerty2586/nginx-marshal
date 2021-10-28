package nginx

import "strings"

func Parse(conf string) Root {
	lines := strings.Split(conf, "\n")
	r := parseSubmodule(lines, "")
	return Root(r)
}

func parseSubmodule(lines []string, name string) Submodule {
	s := Submodule{
		Name:       name,
		Lines:      make([]string, 0),
		Submodules: make([]Submodule, 0),
	}
	for i := 0; i < len(lines); i++ {
		l := rMatchComments.ReplaceAllString(lines[i], "")
		l = strings.TrimSpace(l)
		l = rMatchMultipleSpace.ReplaceAllString(l, " ")
		l = strings.TrimRight(l, ";")
		if len(l) > 0 && l[len(l)-1] == '{' {
			start := i
			brackets := 0
			for ; i < len(lines); i++ {
				brackets += strings.Count(lines[i], "{")
				brackets -= strings.Count(lines[i], "}")
				if brackets == 0 {
					break
				}
			}
			end := i
			name := strings.TrimSpace(strings.TrimRight(l, "{"))
			s.Submodules = append(s.Submodules, parseSubmodule(lines[start+1:end], name))

		} else {
			if l == "" {
				if len(s.Lines) == 0 || s.Lines[len(s.Lines)-1] != "" {
					s.Lines = append(s.Lines, l)
				}
			} else {
				s.Lines = append(s.Lines, l)
			}
		}
	}
	return s
}

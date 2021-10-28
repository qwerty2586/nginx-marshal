package nginx

import (
	"strings"
)

func (s Root) String() string {
	sb := strings.Builder{}
	addLines(&sb, (*Submodule)(&s), 0)
	sb.WriteString("\n")
	addSubmodules(&sb, (*Submodule)(&s), 0)
	return sb.String()
}

func addLines(sb *strings.Builder, s *Submodule, level int) {
	indent := strings.Repeat("    ", level)
	lines := make([]string, len(s.Lines))
	longest_fisrt_word := 0
	for i, l := range s.Lines {
		ll := strings.TrimSpace(l)
		ll = strings.TrimRight(ll, ";")
		if ll == "" {
			lines[i] = ""
			continue
		}
		ll = rMatchMultipleSpace.ReplaceAllString(ll, " ")
		lines[i] = ll
		splits := strings.Split(ll, " ")
		if len(splits) > 0 {
			if len(splits[0]) > longest_fisrt_word {
				longest_fisrt_word = len(splits[0])
			}
		}
	}
	longest_fisrt_word = ((longest_fisrt_word / 4) + 1) * 4
	for _, l := range lines {
		if l == "" {
			sb.WriteString("\n")
			continue
		}
		splits := strings.Split(l, " ")
		if len(splits) > 1 {
			sb.WriteString(indent)
			sb.WriteString(splits[0])
			sb.WriteString(strings.Repeat(" ", longest_fisrt_word+1-len(splits[0])))
			sb.WriteString(strings.Join(splits[1:], " "))
			sb.WriteString(";\n")
		} else {
			sb.WriteString(indent)
			sb.WriteString(l)
			sb.WriteString(";\n")
		}

	}
}

func addSubmodules(sb *strings.Builder, s *Submodule, level int) {
	indent := strings.Repeat("    ", level)
	for _, sub := range s.Submodules {
		sub_name := strings.TrimSpace(sub.Name)

		sb.WriteString(indent)
		sb.WriteString(sub_name)
		sb.WriteString(" {\n")

		addLines(sb, &sub, level+1)
		if len(sub.Submodules) > 0 {
			sb.WriteString("\n")
			addSubmodules(sb, &sub, level+1)
		}
		sb.WriteString(indent)
		sb.WriteString("}\n")
	}
}

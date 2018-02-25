package ox

import (
	"bytes"
	"fmt"
	"strings"
)

type EntryResponse struct {
	Results []Result `json:"results"`
}

func (er *EntryResponse) String() string {
	buf := bytes.NewBufferString("")
	for _, result := range er.Results {
		buf.WriteString(result.String())
	}
	return buf.String()
}

type Result struct {
	Id             string         `json:"id"`
	LexicalEntries []LexicalEntry `json:"lexicalEntries"`
}

func (r *Result) String() string {
	buf := bytes.NewBufferString("")
	buf.WriteString(r.Id)
	buf.WriteString("\n\n")
	for i, v := range r.LexicalEntries {
		buf.WriteString(fmt.Sprintf("%d. Entry: %s\n\n", i+1, v.String()))
	}
	return strings.TrimSpace(buf.String())
}

type LexicalEntry struct {
	Entries  []Entry `json:"entries"`
	Category string  `json:"lexicalCategory"`
}

func (le *LexicalEntry) String() string {
	buf := bytes.NewBufferString("")
	buf.WriteString(le.Category)
	buf.WriteRune('\n')
	for _, entry := range le.Entries {
		buf.WriteString(entry.String())
	}
	return buf.String()
}

type Entry struct {
	Etymologies []string `json:"etymologies"`
	Senses      []Sense  `json:"senses"`
}

func (e *Entry) String() string {
	buf := bytes.NewBufferString("")
	if len(e.Senses) > 0 {
		buf.WriteString("\nSenses:\n")
		for i, v := range e.Senses {
			buf.WriteString(fmt.Sprintf("%c) %s\n", i+'A', v.String()))
		}
	}
	if len(e.Etymologies) > 0 {
		buf.WriteString("\nEtymologies:\n")
		for i, v := range e.Etymologies {
			buf.WriteString(fmt.Sprintf("%c) %s\n", i+'A', v))
		}
	}
	return buf.String()
}

type Feature struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

func (f *Feature) String() string {
	return fmt.Sprintf("%s: %s", f.Type, f.Text)
}

type Sense struct {
	Definitions []string   `json:"definitions"`
	Examples    []Example  `json:"examples"`
	Subsenses   []Subsense `json:"subsenses"`
}

func (s *Sense) String() string {
	buf := bytes.NewBufferString("")
	if len(s.Subsenses) > 0 {
		for _, v := range s.Subsenses {
			buf.WriteString(senseString(v.Definitions, v.Examples))
			buf.WriteString("; ")
		}
	}
	buf.WriteString(senseString(s.Definitions, s.Examples))
	return buf.String()
}

type Example struct {
	Text string `json:"text"`
}

func (e *Example) String() string {
	return e.Text
}

type Subsense struct {
	Definitions []string  `json:"definitions"`
	Examples    []Example `json:"examples"`
}

func (s *Subsense) String() string {
	return senseString(s.Definitions, s.Examples)
}

func senseString(definitions []string, examples []Example) string {
	buf := bytes.NewBufferString("")
	if len(definitions) > 0 {
		buf.WriteString(strings.Join(definitions, "; "))
	}
	if len(examples) > 0 {
		var exampleStrings []string
		buf.WriteString(" (Examples: ")
		for _, v := range examples {
			exampleStrings = append(exampleStrings, v.String())
		}
		buf.WriteString(strings.Join(exampleStrings, "; "))
		buf.WriteRune(')')
	}
	return buf.String()
}

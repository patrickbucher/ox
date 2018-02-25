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
	buf.WriteRune('\n')
	for i, lexicalEntry := range r.LexicalEntries {
		buf.WriteString(fmt.Sprintf("%d. %s", i+1, lexicalEntry.String()))
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
	buf.WriteRune(' ')
	for _, entry := range le.Entries {
		buf.WriteString(entry.String())
	}
	return buf.String()
}

type Entry struct {
	Etymologies []string  `json:"etymologies"`
	Grammar     []Feature `json:"grammaticalFeatures"`
	Senses      []Sense   `json:"senses"`
}

func (e *Entry) String() string {
	buf := bytes.NewBufferString("")
	if len(e.Grammar) > 0 {
		var grammar []string
		for _, feature := range e.Grammar {
			grammar = append(grammar, feature.String())
		}
		buf.WriteRune('(')
		buf.WriteString(strings.Join(grammar, ", "))
		buf.WriteRune(')')
		buf.WriteRune('\n')
	}
	if len(e.Etymologies) > 0 {
		buf.WriteString("Etymologies:\n")
		for i, v := range e.Etymologies {
			buf.WriteString(fmt.Sprintf("%d. %s", i+1, v))
		}
		buf.WriteRune('\n')
	}
	if len(e.Senses) > 0 {
		buf.WriteString("Senses:\n")
		for i, v := range e.Senses {
			buf.WriteString(fmt.Sprintf("%d. %s", i+1, v.String()))
		}
		buf.WriteRune('\n')
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
	// TODO: s.Subsenses
	return senseString(s.Definitions, s.Examples)
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
		buf.WriteString("Meaning: ")
		buf.WriteString(strings.Join(definitions, ";"))
	}
	if len(examples) > 0 {
		var exampleStrings []string
		buf.WriteString(" (Examples: ")
		for _, v := range examples {
			exampleStrings = append(exampleStrings, v.String())
		}
		buf.WriteString(strings.Join(exampleStrings, ";"))
		buf.WriteRune(')')
		buf.WriteRune('\n')
	}
	return buf.String()
}

package ox

type EntryResponse struct {
	Results []Result `json:"results"`
}

type Result struct {
	Id             string         `json:"id"`
	LexicalEntries []LexicalEntry `json:"lexicalEntries"`
}

type LexicalEntry struct {
	Entries  []Entry `json:"entries"`
	Category string  `json:"lexicalCategory"`
}

type Entry struct {
	Etymologies []string  `json:"etymologies"`
	Grammar     []Feature `json:"grammaticalFeatures"`
	Senses      []Sense   `json:"senses"`
}

type Feature struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type Sense struct {
	Definitions []string   `json:"definitions"`
	Examples    []Example  `json:"examples"`
	Subsenses   []Subsense `json:"subsenses"`
}

type Example struct {
	Text string `json:"text"`
}

type Subsense struct {
	Definitions []string  `json:"definitions"`
	Examples    []Example `json:"examples"`
}

package ent

import "strings"

func NewTermQuery(field, keyword string) *TermQuery {
	return &TermQuery{
		Keyword: &Keyword{
			Field: field,
			Word:  keyword,
		},
	}
}

func (t *TermQuery) Empty() bool {
	return t.Keyword == nil && len(t.Must) == 0 && len(t.Should) == 0
}

func (t *TermQuery) And(queries ...*TermQuery) *TermQuery {
	if len(queries) == 0 {
		return t
	}

	musts := make([]*TermQuery, 0, len(queries)+1)

	if !t.Empty() {
		musts = append(musts, t)
	}

	for _, query := range queries {
		if !query.Empty() {
			musts = append(musts, query)
		}
	}

	return &TermQuery{Must: musts}
}

func (t *TermQuery) Or(queries ...*TermQuery) *TermQuery {
	if len(queries) == 0 {
		return t
	}

	should := make([]*TermQuery, 0, len(queries)+1)

	if !t.Empty() {
		should = append(should, t)
	}

	for _, query := range queries {
		if !query.Empty() {
			should = append(should, query)
		}
	}

	return &TermQuery{Should: should}
}

func (t *TermQuery) Build() string {
	// if keyword not empty, means this TermQuery is the root node
	if t.Keyword != nil {
		return t.Keyword.IntoString()
	}

	numMust := len(t.Must)

	if numMust > 0 {
		if numMust == 1 {
			return t.Must[0].Build()
		}
		b := strings.Builder{}
		b.WriteByte('(')

		for i, m := range t.Must {
			t := m.Build()
			if len(t) > 0 {
				b.WriteString(t)
				if i < numMust-1 {
					b.WriteByte('&')
				}
			}
		}

		b.WriteByte(')')

		return b.String()
	}

	numShould := len(t.Should)

	if numShould > 0 {
		if numShould == 1 {
			return t.Should[0].Build()
		}
		b := strings.Builder{}
		b.WriteByte('(')

		for i, s := range t.Should {
			t := s.Build()
			if len(t) > 0 {
				b.WriteString(t)
				if i < numMust-1 {
					b.WriteByte('|')
				}
			}
		}

		b.WriteByte(')')

		return b.String()
	}

	return ""
}

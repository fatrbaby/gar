package ent

import (
	"testing"
)

func TestTermQuery_Build(t *testing.T) {
	r := NewTermQuery("name", "fat").
		And(NewTermQuery("name", "baby")).
		And(NewTermQuery("name", "cute")).Build()
	t.Log(r)
}

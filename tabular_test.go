package tabular

import (
	"reflect"
	"testing"
)

func TestTabular(t *testing.T) {
	want := Format{
		"env": &Column{Name: "Environment", Length: 14},
		"cls": &Column{Name: "Cluster", Length: 10},
		"svc": &Column{Name: "Service", Length: 25},
		"hst": &Column{Name: "Database Host", Length: 25},
	}

	got := New()
	for k, v := range want {
		got.Add(k, v.Name, v.Length)
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("ERROR: tabular.Add() failed\n  want: %#v\n   got: %#v", got, want)
	}
}

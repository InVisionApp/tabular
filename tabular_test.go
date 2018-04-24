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
		"pct": &Column{Name: "PCT", Length: 5, RightJustified: true},
	}

	got := New()
	for k, v := range want {
		got.Add(k, v.Name, v.Length)
		got[k].RightJustified = v.RightJustified
	}

	if !reflect.DeepEqual(want, got) {
		if len(got) != len(want) {
			t.Fatalf("ERROR: tabular.Add() failed to produce map of the right size")
		}
		for k, v := range want {
			if !reflect.DeepEqual(v, got[k]) {
				t.Errorf("ERROR: tabular.Column mismatch\n  want: %#v\n   got: %#v", *v, *got[k])
			}
		}

	}
}

package tabular

import (
	"reflect"
	"testing"
)

func TestTabular(t *testing.T) {
	want := Format{
		"env": Column{n: "Environment", l: 14},
		"cls": Column{n: "Cluster", l: 10},
		"svc": Column{n: "Service", l: 25},
		"hst": Column{n: "Database Host", l: 25},
	}

	got := New()
	for k, v := range want {
		got.Add(k, v.n, v.l)
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("ERROR: tabular.Add() failed\n  want: %#v\n   got: %#v", got, want)
	}
}

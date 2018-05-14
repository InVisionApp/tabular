package tabular

import (
	"reflect"
	"testing"
)

func TestTabular(t *testing.T) {
	want := Table{
		Columns: map[string]*Column{
			"env": &Column{Name: "Environment", Length: 14},
			"cls": &Column{Name: "Cluster", Length: 10},
			"svc": &Column{Name: "Service", Length: 25},
			"hst": &Column{Name: "Database Host", Length: 25},
			"pct": &Column{Name: "PCT", Length: 5, RightJustified: true},
		},
		order: &[]string{
			"env",
			"cls",
			"svc",
			"hst",
			"pct",
		},
	}

	got := New()
	for _, v := range *want.order {
		c := want.Columns[v]
		f := got.Col
		if c.RightJustified {
			f = got.ColRJ
		}
		f(v, c.Name, c.Length)
	}

	if !reflect.DeepEqual(want, got) {
		if len(got.Columns) != len(want.Columns) {
			t.Fatalf("ERROR: tabular failed to produce map of the right size\nWant: %d, Got: %d\n", len(want.Columns), len(got.Columns))
		}
		if len(*got.order) != len(*want.order) {
			t.Fatalf("ERROR: tabular failed to produce array of the right size\nWant: %d, Got: %d\n", len(*want.order), len(*got.order))
		}
		for k, v := range want.Columns {
			if !reflect.DeepEqual(v, got.Columns[k]) {
				t.Errorf("ERROR: tabular.Column mismatch\n  want: %#v\n   got: %#v", *v, *got.Columns[k])
			}
		}
		for k, v := range *want.order {
			if v != (*got.order)[k] {
				t.Errorf("ERROR: tabular.order mismatch\n  want: %#v\n   got: %#v", v, (*got.order)[k])
			}
		}

	}
}

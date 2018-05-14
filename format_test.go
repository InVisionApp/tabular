package tabular_test

import (
	"testing"

	"github.com/InVisionApp/tabular"
)

func TestFormat(t *testing.T) {
	tab := tabular.New()
	tab.ColRJ("id", "ID", 6)
	tab.Col("env", "Environment", 14)
	tab.Col("cls", "Cluster", 10)
	tab.Col("svc", "Service", 25)
	tab.Col("hst", "Database Host", 25)
	tab.ColRJ("pct", "%CPU", 5)

	// Test Partial Printing
	want := "%6v %-14v %-10v\n"
	if got := tab.Print("id", "env", "cls"); got != want {
		t.Errorf("ERROR: tab.Print() failed\n   want: %q\n    got: %q", want, got)
	}

	tWant := tabular.Output{
		Header:    "    ID Environment    Cluster    Service                   Database Host              %CPU",
		SubHeader: "------ -------------- ---------- ------------------------- ------------------------- -----",
		Format:    "%6v %-14v %-10v %-25v %-25v %5v\n",
	}

	// Test Printing All
	want = tWant.Format
	if got := tab.Print(tabular.All); got != want {
		t.Errorf("ERROR: tab.Print(All) failed\n   want: %q\n    got: %q", want, got)
	}

	// Test Parsing
	if tGot := tab.Parse("id", "env", "cls", "svc", "hst", "pct"); tGot != tWant {
		if tGot.Header != tWant.Header {
			t.Errorf("ERROR: tab.Parse() failed\n   want: %q\n    got: %q", tWant.Header, tGot.Header)
		}
		if tGot.SubHeader != tWant.SubHeader {
			t.Errorf("ERROR: tab.Parse() failed\n   want: %q\n    got: %q", tWant.SubHeader, tGot.SubHeader)
		}
		if tGot.Format != tWant.Format {
			t.Errorf("ERROR: tab.Parse() failed\n   want: %q\n    got: %q", tWant.Format, tGot.Format)
		}
	}
}

package tabular_test

import (
	"testing"

	"github.com/InVisionApp/tabular"
)

func TestFormat(t *testing.T) {
	tab := tabular.New()
	tab.Col("id", "ID", 6)
	tab.Col("env", "Environment", 14)
	tab.Col("cls", "Cluster", 10)
	tab.Col("svc", "Service", 25)
	tab.Col("hst", "Database Host", 25)
	tab.Col("pct", "%CPU", 5)
	tab["id"].RightJustified = true
	tab["pct"].RightJustified = true

	tWant := tabular.Table{
		Header:    "    ID Environment    Cluster    Service                   Database Host              %CPU",
		SubHeader: "------ -------------- ---------- ------------------------- ------------------------- -----",
		Format:    "%6v %-14v %-10v %-25v %-25v %5v\n",
	}

	// Test Printing
	want := tWant.Format
	if got := tab.Print("id", "env", "cls", "svc", "hst", "pct"); got != want {
		t.Errorf("ERROR: tab.Print() failed\n   want: %q\n    got: %q", want, got)
	}

	// Test Parsing
	if tGot := tab.Parse("id", "env", "cls", "svc", "hst", "pct"); tGot != tWant {
		if tGot.Header != tWant.Header {
			t.Errorf("ERROR: tab.Parse() failed\n   want: %v\n    got: %v", tWant.Header, tGot.Header)
		}
		if tGot.SubHeader != tWant.SubHeader {
			t.Errorf("ERROR: tab.Parse() failed\n   want: %v\n    got: %v", tWant.SubHeader, tGot.SubHeader)
		}
		if tGot.Format != tWant.Format {
			t.Errorf("ERROR: tab.Parse() failed\n   want: %v\n    got: %v", tWant.Format, tGot.Format)
		}
	}
}

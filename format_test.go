package tabular_test

import (
	"testing"

	"github.com/InVisionApp/tabular"
)

func TestFormat(t *testing.T) {
	tab := tabular.New()
	tab.Col("env", "Environment", 14)
	tab.Col("cls", "Cluster", 10)
	tab.Col("svc", "Service", 25)
	tab.Col("hst", "Database Host", 25)
	tab.Col("pct", "%CPU", 5)
	tab["pct"].RightJustified = true

	tWant := tabular.Table{
		Header:    "Environment    Cluster    Service                   Database Host              %CPU",
		SubHeader: "-------------- ---------- ------------------------- ------------------------- -----",
		Format:    "%-14v %-10v %-25v %-25v %5v\n",
	}

	// Test Printing
	want := tWant.Format
	if got := tab.Print("env", "cls", "svc", "hst", "pct"); got != want {
		t.Errorf("ERROR: tab.Print() failed\n   want: %q\n    got: %q", want, got)
	}

	// Test Parsing
	if tGot := tab.Parse("env", "cls", "svc", "hst", "pct"); tGot != tWant {
		t.Errorf("ERROR: tab.Parse() failed\n   want: %v\n    got: %v", tWant, tGot)
	}
}

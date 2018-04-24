package tabular_test

import (
	"testing"

	"github.com/InVisionApp/tabular"
)

func TestFormat(t *testing.T) {
	tab := tabular.New()
	tab.Add("env", "Environment", 14)
	tab.Add("cls", "Cluster", 10)
	tab.Add("svc", "Service", 25)
	tab.Add("hst", "Database Host", 25)
	tab.Add("pct", "%CPU", 5)
	tab["pct"].RightJustified = true

	want := "%-14v %-10v %-25v %-25v %5v\n"
	if got := tab.Do("env", "cls", "svc", "hst", "pct"); got != want {
		t.Fatalf("ERROR: tab.Do() failed\n   want: %q\n    got: %q", want, got)
	}
}

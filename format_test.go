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

	want := "%-14s %-10s %-25s %-25s\n"
	if got := tab.Do("env", "cls", "svc", "hst"); got != want {
		t.Fatalf("ERROR: tab.Do() failed\n   want: %q\n    got: %q", want, got)
	}
}

package tabular

import (
	"fmt"
	"strings"
)

// Format - maps short names of columns to their structure defining:
//	full name, length and whether it's right justified
//
// For Example:
// 	"env": Column{Name: "Environment", Length: 14},
// 	"cls": Column{Name: "Cluster",     Length: 40},
// 	"srv": Column{Name: "Service",     Length: 35},
// 	"hst": Column{Name: "Host",        Length: 45},
// 	"pct": Column{Name: "%CPU",        Length: 7, RightJustified: true},
type Format map[string]*Column

// Column - defines column's name, length and if it's right justified
type Column struct {
	Name           string
	Length         int
	RightJustified bool
}

// Do - does the following:
//
// 1) prints a table style heading for a given list of columns.
//
//    For example if Format is defined as:
//
//      "env": Column{Name: "Environment", Length: 14},
//      "cls": Column{Name: "Cluster",     Length: 40},
//      "srv": Column{Name: "Service",     Length: 35},
//
//    It'll produce:
//
// 	Environment    Cluster                                  Service
//	-------------- ---------------------------------------- -----------------------------------
//
// 2) Returns an fmt style `format` string to output values
//    under the above heading via Printf(format,...):
//
//	%-14v %-40v %-35v
func (fm Format) Do(cols ...string) string {
	var title string
	var uline string
	var format string
	for _, c := range cols {
		title = title + " " + fmt.Sprintf(fm[c].f(), fm[c].Name)
		uline = uline + " " + fmt.Sprintf(fm[c].f(), r(fm[c].Length))
		format = format + " " + fm[c].f()
	}
	fmt.Println(strings.TrimSpace(title))
	fmt.Println(strings.TrimSpace(uline))
	return strings.TrimSpace(format) + "\n"
}

// New - Creates a new tabular Format
func New() Format { return Format{} }

// Add - adds a new column to existing tabular Format
func (fm Format) Add(shortName, fullName string, columnLength int) {
	fm[shortName] = &Column{Name: fullName, Length: columnLength}
}

// f() returns fmt formatting, for example:
//
// 	Column{Name: "Environment", Length: 14, RightJustified: false}
// 	result => %-14v
//
// 	Column{Name: "PCT", Length: 5, RightJustified: true}
// 	result => %5v
func (c *Column) f() string {
	pad := "-"
	if c.RightJustified {
		pad = ""
	}
	return fmt.Sprintf("%%%s%dv", pad, c.Length)
}

// r() returns a dashed line for table formatting "------"
// with it's length set to the length of l
func r(l int) string { return strings.Repeat("-", l) }

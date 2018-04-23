package tabular

import (
	"fmt"
	"strings"
)

// Format - maps short names of columns to a structure defining their full names and lengths
//
// For Example:
// 	"env": Column{n: "Environment", l: 14},
// 	"cls": Column{n: "Cluster", l: 40},
// 	"srv": Column{n: "Service", l: 35},
// 	"cr8": Column{n: "CreatedAt", l: 19},
// 	"vld": Column{n: "Valid", l: 5},
// 	"dbt": Column{n: "DBType", l: 10},
// 	"hst": Column{n: "Host", l: 45},
type Format map[string]Column

// Column - defines column's name and length
type Column struct {
	n string
	l int
}

// Do - does the following:
//
// 1) prints a table style heading for a given list of columns.
//
//    For example if Format is defined as:
//
// 	"env": Column{n: "Environment", l: 14},
//      "cls": Column{n: "Cluster", l: 40},
//      "srv": Column{n: "Service", l: 35},
//
//    It'll produce:
//
// 	Environment    Cluster                                  Service
//	-------------- ---------------------------------------- -----------------------------------
//
// 2) Returns an fmt style format string to push values into this table:
//
//	%-14s %-40s %-35s
func (fm Format) Do(cols ...string) string {
	var title string
	var uline string
	var format string
	for _, c := range cols {
		title = title + " " + fmt.Sprintf(fm[c].f(), fm[c].n)
		uline = uline + " " + fmt.Sprintf(fm[c].f(), r(fm[c].l))
		format = format + " " + fm[c].f()
	}
	fmt.Println(strings.TrimSpace(title))
	fmt.Println(strings.TrimSpace(uline))
	return strings.TrimSpace(format) + "\n"
}

// New - Creates a new tabular Format
func New() Format {
	return Format{}
}

// Add - adds a new column to existing tabular Format
func (fm Format) Add(shortName, fullName string, columnLength int) {
	fm[shortName] = Column{n: fullName, l: columnLength}
}

// f() returns fmt formatting in a form of `%-14s`
// for example if a Column is defined as:
// 	Column{n: "Environment", l: 14}
// }
// It'll return: %-14s
func (c Column) f() string {
	return fmt.Sprintf("%%-%ds", c.l)
}

// r() returns a underbar line for table formatting "------"
// with it's length set to the length of c
func r(l int) string { return strings.Repeat("-", l) }

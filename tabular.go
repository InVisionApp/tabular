package tabular

import (
	"fmt"
	"strings"
)

// Table - parsed table's header, subheader and format specifier
type Table struct {
	Header    string
	SubHeader string
	Format    string
}

// Columns - maps short names of columns to their structure defining:
//	full name, length and whether it's right justified
//
// For Example:
// 	"env": Column{Name: "Environment", Length: 14},
// 	"cls": Column{Name: "Cluster",     Length: 40},
// 	"srv": Column{Name: "Service",     Length: 35},
// 	"hst": Column{Name: "Host",        Length: 45},
// 	"pct": Column{Name: "%CPU",        Length: 7, RightJustified: true},
type Columns map[string]*Column

// Column - defines column's name, length and if it's right justified
type Column struct {
	Name           string
	Length         int
	RightJustified bool
}

// New - Creates a map of tabular Columns
func New() Columns { return Columns{} }

// Print - does the following:
//
// 1) prints a table style heading for a given list of columns.
//
//    For example if Columns are defined as:
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
// 2) Returns an fmt style format specifier string that you can use
//    to output values under the above heading via Printf(format,...):
//
//	%-14v %-40v %-35v
func (cl Columns) Print(cols ...string) string {
	t := cl.parse(cols...)
	fmt.Println(t.Header)
	fmt.Println(t.SubHeader)
	return t.Format
}

// Parse - builds a Table out of a given list of columns
//
// To simply print the table's title call Print() instead
//
// Parse() is usefull when you need to control where
// to output the title, for example to a log or a trace file
func (cl Columns) Parse(cols ...string) Table {
	return cl.parse(cols...)
}

// Col - adds a new column to existing tabular Format
func (cl Columns) Col(shortName, fullName string, columnLength int) {
	cl[shortName] = &Column{Name: fullName, Length: columnLength}
}

func (cl Columns) parse(cols ...string) Table {
	var header string
	var subHeader string
	var format string
	for _, c := range cols {
		header = header + " " + fmt.Sprintf(cl[c].f(), cl[c].Name)
		subHeader = subHeader + " " + fmt.Sprintf(cl[c].f(), r(cl[c].Length))
		format = format + " " + cl[c].f()
	}

	return Table{
		Header:    strings.TrimSpace(header),
		SubHeader: strings.TrimSpace(subHeader),
		Format:    strings.TrimSpace(format) + "\n",
	}
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

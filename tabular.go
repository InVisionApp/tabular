package tabular

import (
	"fmt"
	"strings"
)

// Output - parsed table's header, subheader and format specifier
type Output struct {
	Header    string
	SubHeader string
	Format    string
}

// Table - maps and orders short names of columns to their structure defining:
// full name, length and whether it's right justified
type Table struct {
	Columns map[string]*Column
	order   *[]string
}

// Column - defines column's name, length and if it's right justified
type Column struct {
	Name           string
	Length         int
	RightJustified bool
}

// All - pass this to Print() or Parse() to print or parse all columns of a table
const All = "*"

// New - Creates a new table
func New() Table {
	return Table{
		Columns: map[string]*Column{},
		order:   &[]string{},
	}
}

// Print - does the following:
//
// 1) prints a table style heading for a given list of columns,
// for example, if Columns are defined as:
//
//      "env": Column{Name: "Environment", Length: 14},
//      "cls": Column{Name: "Cluster",     Length: 40},
//      "srv": Column{Name: "Service",     Length: 35},
//
// It'll produce:
//
// 	Environment    Cluster                                  Service
//	-------------- ---------------------------------------- -----------------------------------
//
// 2) Returns an fmt style format specifier string that you can use
// to output values under the above heading via Printf(format,...):
//
//	%-14v %-40v %-35v
func (tbl Table) Print(cols ...string) string {
	t := tbl.parse(cols...)
	fmt.Println(t.Header)
	fmt.Println(t.SubHeader)
	return t.Format
}

// Parse - constructs Table's Output structure containing it's header,
// sub-header and format modifier out of a given list of columns.
//
// To simply print the table's title call Print() instead.
//
// Parse() is usefull when you need to control where
// to output the title, for example to a log or a trace file.
func (tbl Table) Parse(cols ...string) Output {
	return tbl.parse(cols...)
}

// Col - adds a new column to an existing table
func (tbl Table) Col(shortName, fullName string, columnLength int) {
	tbl.Columns[shortName] = &Column{Name: fullName, Length: columnLength}
	tbl.appendColumn(shortName)
}

// ColRJ - adds a new Right Justified column to an existing table
func (tbl Table) ColRJ(shortName, fullName string, columnLength int) {
	tbl.Columns[shortName] = &Column{Name: fullName, Length: columnLength, RightJustified: true}
	tbl.appendColumn(shortName)
}

func (tbl Table) appendColumn(shortName string) {
	*tbl.order = append(*tbl.order, shortName)
}

func (tbl Table) parse(cols ...string) Output {
	var header string
	var subHeader string
	var format string
	var space string

	if len(cols) == 1 && cols[0] == All {
		cols = *tbl.order
	}

	for _, c := range cols {
		cf := tbl.Columns[c].f()
		header = header + space + fmt.Sprintf(cf, tbl.Columns[c].Name)
		subHeader = subHeader + space + fmt.Sprintf(cf, r(tbl.Columns[c].Length))
		format = format + space + cf
		space = " "
	}

	return Output{
		Header:    header,
		SubHeader: subHeader,
		Format:    format + "\n",
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

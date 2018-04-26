[![LICENSE](https://img.shields.io/badge/license-MIT-orange.svg)](LICENSE)
[![Build Status](https://travis-ci.org/InVisionApp/tabular.svg?branch=master)](https://travis-ci.org/InVisionApp/tabular)
[![codecov](https://codecov.io/gh/InVisionApp/tabular/branch/master/graph/badge.svg)](https://codecov.io/gh/InVisionApp/tabular)
[![Go Report Card](https://goreportcard.com/badge/github.com/InVisionApp/tabular)](https://goreportcard.com/report/github.com/InVisionApp/tabular)
[![Godocs](https://img.shields.io/badge/golang-documentation-blue.svg)](https://godoc.org/github.com/InVisionApp/tabular)

# tabular

Tabular simplifies printing ASCII tables from command line utilities without the need to pass large sets of data to it's API.  

Simply define the table columns and `tabular` will parse the right [format specifier](https://golang.org/pkg/fmt/#Printf) that you can use in your calls to `fmt.Printf()` or any other function that supports it.

Table columns can be defined once and then reused over and over again making it easy to modify column length and heading in one place.  And a subset of columns can be specified during `tabular.Print()` or `tabular.Parse()` calls to modify the table's title without redefining it.

Example (also available in [`example/example.go`](example/example.go):

```go
package main

import (
	"fmt"
	"log"

	"github.com/InVisionApp/tabular"
)

var tab tabular.Columns

func init() {
	tab = tabular.New()
	tab.Col("env", "Environment", 14)
	tab.Col("cls", "Cluster", 10)
	tab.Col("svc", "Service", 15)
	tab.Col("hst", "Database Host", 20)
	tab.Col("pct", "%CPU", 7)
	tab["pct"].RightJustified = true
}

// Sample data-set
var data = []struct {
	e, c, s, d string
	v          float64
}{
	{
		e: "production",
		c: "cluster-1",
		s: "service-a",
		d: "database-host-1",
		v: 70.01,
	},
	{
		e: "production",
		c: "cluster-1",
		s: "service-b",
		d: "database-host-2",
		v: 99.51,
	},
	{
		e: "production",
		c: "cluster-2",
		s: "service-a",
		d: "database-host-1",
		v: 70.01,
	},
	{
		e: "production",
		c: "cluster-2",
		s: "service-b",
		d: "database-host-2",
		v: 99.51,
	},
}

func main() {
	// Print Environments and Clusters
	format := tab.Print("env", "cls")
	for _, x := range data {
		fmt.Printf(format, x.e, x.c)
	}

	// Print Environments, Clusters and Services
	format = tab.Print("env", "cls", "svc")
	for _, x := range data {
		fmt.Printf(format, x.e, x.c, x.s)
	}

	// Print Everything
	format = tab.Print("cls", "svc", "hst", "pct")
	for _, x := range data {
		fmt.Printf(format, x.c, x.s, x.d, x.v)
	}

	// Print Everything to a custom destination such as a log
	table := tab.Parse("cls", "svc", "hst", "pct")
	log.Println(table.Header)
	log.Println(table.SubHeader)
	for _, x := range data {
		log.Printf(table.Format, x.c, x.s, x.d, x.v)
	}
}
```

Produces:

```
Environment    Cluster
-------------- ----------
production     cluster-1
production     cluster-1
production     cluster-2
production     cluster-2

Environment    Cluster    Service
-------------- ---------- ---------------
production     cluster-1  service-a
production     cluster-1  service-b
production     cluster-2  service-a
production     cluster-2  service-b

Cluster    Service         Database Host           %CPU
---------- --------------- -------------------- -------
cluster-1  service-a       database-host-1        70.01
cluster-1  service-b       database-host-2        99.51
cluster-2  service-a       database-host-1        70.01
cluster-2  service-b       database-host-2        99.51

2018/04/26 10:17:27 Cluster    Service         Database Host           %CPU
2018/04/26 10:17:27 ---------- --------------- -------------------- -------
2018/04/26 10:17:27 cluster-1  service-a       database-host-1        70.01
2018/04/26 10:17:27 cluster-1  service-b       database-host-2        99.51
2018/04/26 10:17:27 cluster-2  service-a       database-host-1        70.01
2018/04/26 10:17:27 cluster-2  service-b       database-host-2        99.51
```
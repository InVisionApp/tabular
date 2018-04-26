[![LICENSE](https://img.shields.io/badge/license-MIT-orange.svg)](LICENSE)
[![Build Status](https://travis-ci.org/InVisionApp/tabular.svg?branch=master)](https://travis-ci.org/InVisionApp/tabular)
[![codecov](https://codecov.io/gh/InVisionApp/tabular/branch/master/graph/badge.svg)](https://codecov.io/gh/InVisionApp/tabular)
[![Go Report Card](https://goreportcard.com/badge/github.com/InVisionApp/tabular)](https://goreportcard.com/report/github.com/InVisionApp/tabular)
[![Godocs](https://img.shields.io/badge/golang-documentation-blue.svg)](https://godoc.org/github.com/InVisionApp/tabular)

# tabular

Tabular package to print ASCII tables from command line utilities.

Example:

```go
package main

import (
	"fmt"

	"github.com/InVisionApp/tabular"
)

var tab tabular.Format

func init() {
	tab = tabular.New()
	tab.Add("env", "Environment", 14)
	tab.Add("cls", "Cluster", 10)
	tab.Add("svc", "Service", 15)
	tab.Add("hst", "Database Host", 20)
	tab.Add("pct", "%CPU", 7)
	tab["pct"].RightJustified = true
}

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
	format := tab.Do("env", "cls")
	for _, x := range data {
		fmt.Printf(format, x.e, x.c)
	}

	// Print Environments, Clusters and Services
	format = tab.Do("env", "cls", "svc")
	for _, x := range data {
		fmt.Printf(format, x.e, x.c, x.s)
	}

	// Print Clusters, Services and Database Hosts
	format = tab.Do("cls", "svc", "hst", "pct")
	for _, x := range data {
		fmt.Printf(format, x.c, x.s, x.d, x.v)
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
```
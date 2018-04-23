[![LICENSE](https://img.shields.io/badge/license-MIT-orange.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/InVisionApp/tabular)](https://goreportcard.com/report/github.com/InVisionApp/tabular)
[![Godocs](https://img.shields.io/badge/golang-documentation-blue.svg)](https://godoc.org/github.com/InVisionApp/tabular)

# tabular

Tabular package to print ASCII tables from command line utilities.

Example:

```
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
	tab.Add("svc", "Service", 25)
	tab.Add("hst", "Database Host", 25)
}

var data = []struct {
	e, c, s, d string
}{
	{
		e: "production",
		c: "cluster-1",
		s: "service-a",
		d: "database-host-1",
	},
	{
		e: "production",
		c: "cluster-1",
		s: "service-b",
		d: "database-host-2",
	},
	{
		e: "production",
		c: "cluster-2",
		s: "service-a",
		d: "database-host-1",
	},
	{
		e: "production",
		c: "cluster-2",
		s: "service-b",
		d: "database-host-2",
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
	format = tab.Do("cls", "svc", "hst")
	for _, x := range data {
		fmt.Printf(format, x.c, x.s, x.d)
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
-------------- ---------- -------------------------
production     cluster-1  service-a
production     cluster-1  service-b
production     cluster-2  service-a
production     cluster-2  service-b

Cluster    Service                   Database Host
---------- ------------------------- -------------------------
cluster-1  service-a                 database-host-1
cluster-1  service-b                 database-host-2
cluster-2  service-a                 database-host-1
cluster-2  service-b                 database-host-2
```
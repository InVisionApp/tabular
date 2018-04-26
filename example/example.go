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

	// Print Clusters, Services and Database Hosts
	format = tab.Print("cls", "svc", "hst", "pct")
	for _, x := range data {
		fmt.Printf(format, x.c, x.s, x.d, x.v)
	}

	// Print to a custom destination such as a log
	table := tab.Parse("cls", "svc", "hst", "pct")
	log.Println(table.Header)
	log.Println(table.SubHeader)
	for _, x := range data {
		log.Printf(table.Format, x.c, x.s, x.d, x.v)
	}
}

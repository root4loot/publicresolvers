package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/root4loot/goresolvers"
)

func banner() {
	fmt.Println("\nA tool for fetching DNS resolver lists from the Trickest Resolvers repository (https://github.com/trickest/resolvers)")
	fmt.Println("")
}

func usage() {
	// create a new tabwriter
	w := tabwriter.NewWriter(os.Stdout, 2, 0, 3, ' ', 0)

	// print the usage header
	fmt.Fprintln(w, "Usage: "+os.Args[0]+" [flag]\n")

	fmt.Fprintf(w, "  %s,\t%s\t%s\t(%s)\n", "-r", "--resolvers", "fetch resolvers.txt", "resolver IP addresses")
	fmt.Fprintf(w, "  %s,\t%s\t%s\t(%s)\n", "-t", "--trusted", "fetch resolvers-trusted.txt", "trusted resolvers from organizations like Cloudflare, Google, etc.")
	fmt.Fprintf(w, "  %s,\t%s\t%s\t(%s)\n", "-c", "--community", "fetch resolvers-community.txt", "resolver IP addresses with community annotations")
	fmt.Fprintf(w, "  %s,\t%s\t%s\n", "-h", "--help", "display help")

	// flush the tabwriter
	w.Flush()
}

func printList(source func() ([]string, error), name string) {
	list, err := source()
	if err != nil {
		fmt.Printf("Error fetching %s resolvers: %v\n", name, err)
		return
	}
	fmt.Printf("%s resolvers:\n%s\n", name, strings.Join(list, "\n"))
}

func main() {
	var resolvers, trusted, community, help bool

	flag.BoolVar(&resolvers, "r", false, "")
	flag.BoolVar(&resolvers, "resolvers", false, "")
	flag.BoolVar(&trusted, "t", false, "")
	flag.BoolVar(&trusted, "trusted", false, "")
	flag.BoolVar(&community, "c", false, "")
	flag.BoolVar(&community, "community", false, "")
	flag.BoolVar(&help, "h", false, "")
	flag.BoolVar(&help, "help", false, "")
	flag.Parse()

	if help {
		banner()
		usage()
		os.Exit(0)
	}

	if !resolvers && !community && !trusted {
		fmt.Println("Error: No flags provided.")
		usage()
		os.Exit(1)
	}

	if resolvers {
		printList(goresolvers.FetchResolvers, "Resolvers")
	}

	if trusted {
		printList(goresolvers.FetchResolversTrusted, "Trusted")
	}

	if community {
		printList(goresolvers.FetchResolversCommunity, "Community")
	}
}

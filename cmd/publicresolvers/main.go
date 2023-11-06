package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/root4loot/publicresolvers"
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
	fmt.Fprintf(w, "  %s,\t%s\t%s\n", "-p", "--with-port", "include port 53 in resolver IP addresses")
	fmt.Fprintf(w, "  %s,\t%s\t%s\n", "-h", "--help", "display help")

	// flush the tabwriter
	w.Flush()
}

func printList(source func() ([]string, error)) {
	list, err := source()
	if err != nil {
		fmt.Printf("Error fetching resolvers: %v\n", err)
		return
	}
	fmt.Printf("%s\n", strings.Join(list, "\n"))
}

func main() {
	var resolvers, trusted, community, help, withPort bool

	flag.BoolVar(&resolvers, "r", false, "")
	flag.BoolVar(&resolvers, "resolvers", false, "")
	flag.BoolVar(&trusted, "t", false, "")
	flag.BoolVar(&trusted, "trusted", false, "")
	flag.BoolVar(&community, "c", false, "")
	flag.BoolVar(&community, "community", false, "")
	flag.BoolVar(&withPort, "p", false, "")
	flag.BoolVar(&withPort, "with-port", false, "")
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
		if !withPort {
			printList(publicresolvers.FetchResolvers)
		} else {
			printList(publicresolvers.FetchResolversWithPort)
		}
	}

	if trusted {
		if !withPort {
			printList(publicresolvers.FetchResolversTrusted)
		} else {
			printList(publicresolvers.FetchResolversTrustedWithPort)
		}
	}

	if community {
		if !withPort {
			printList(publicresolvers.FetchResolversCommunity)
		} else {
			printList(publicresolvers.FetchResolversCommunityWithPort)
		}
	}

}

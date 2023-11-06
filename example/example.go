package main

import (
	"fmt"

	"github.com/root4loot/publicresolvers"
)

func main() {
	resolvers, err := publicresolvers.FetchResolvers()
	if err != nil {
		fmt.Println("Error fetching resolvers:", err)
		return
	}
	fmt.Println(resolvers)

	trustedResolvers, err := publicresolvers.FetchResolversTrusted()
	if err != nil {
		fmt.Println("Error fetching trusted resolvers:", err)
		return
	}
	fmt.Println(trustedResolvers)

	communityResolvers, err := publicresolvers.FetchResolversCommunity()
	if err != nil {
		fmt.Println("Error fetching community resolvers:", err)
		return
	}
	fmt.Println(communityResolvers)
}

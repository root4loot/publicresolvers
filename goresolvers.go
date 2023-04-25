package goresolvers

import (
	"bufio"
	"net/http"
)

const (
	resolversURL          = "https://raw.githubusercontent.com/trickest/resolvers/main/resolvers.txt"
	resolversTrustedURL   = "https://raw.githubusercontent.com/trickest/resolvers/main/resolvers-trusted.txt"
	resolversCommunityURL = "https://raw.githubusercontent.com/trickest/resolvers/main/resolvers-community.txt"
)

func fetchFile(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// FetchResolvers fetches the list of resolvers
func FetchResolvers() ([]string, error) {
	return fetchFile(resolversURL)
}

// FetchResolversTrusted fetches the list of trusted resolvers
func FetchResolversTrusted() ([]string, error) {
	return fetchFile(resolversTrustedURL)
}

func FetchResolversCommunity() ([]string, error) {
	return fetchFile(resolversCommunityURL)
}

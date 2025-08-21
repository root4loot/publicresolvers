package publicresolvers

import (
	"bufio"
	"net/http"
)

const (
	resolversURL          = "https://raw.githubusercontent.com/trickest/resolvers/refs/heads/main/resolvers.txt"
	resolversTrustedURL   = "https://raw.githubusercontent.com/trickest/resolvers/refs/heads/main/resolvers-trusted.txt"
	resolversCommunityURL = "https://raw.githubusercontent.com/trickest/resolvers/refs/heads/main/resolvers-community.txt"
)

func fetchFile(url string, includePort bool) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		if includePort {
			line += ":53"
		}
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// FetchResolvers fetches the list of resolvers from the Trickest Resolvers repository.
func FetchResolvers() ([]string, error) {
	return fetchFile(resolversURL, false)
}

// FetchResolversWithPort fetches the list of resolvers from the Trickest Resolvers repository with port 53 appended.
func FetchResolversWithPort() ([]string, error) {
	return fetchFile(resolversURL, true)
}

// FetchResolversTrusted fetches the list of trusted resolvers from the Trickest Resolvers repository.
func FetchResolversTrusted() ([]string, error) {
	return fetchFile(resolversTrustedURL, false)
}

// FetchResolversTrustedWithPort fetches the list of trusted resolvers from the Trickest Resolvers repository with port 53 appended.
func FetchResolversTrustedWithPort() ([]string, error) {
	return fetchFile(resolversTrustedURL, true)
}

// FetchResolversCommunity fetches the list of community resolvers from the Trickest Resolvers repository.
func FetchResolversCommunity() ([]string, error) {
	return fetchFile(resolversCommunityURL, false)
}

// FetchResolversCommunityWithPort fetches the list of community resolvers from the Trickest Resolvers repository with port 53 appended.
func FetchResolversCommunityWithPort() ([]string, error) {
	return fetchFile(resolversCommunityURL, true)
}

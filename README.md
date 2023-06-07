![Go version](https://img.shields.io/badge/Go-v1.19-blue.svg) [![Contribute](https://img.shields.io/badge/Contribute-Welcome-green.svg)](CONTRIBUTING.md)

# publicresolvers
publicresolvers is a simple Go library and CLI tool for fetching DNS resolver lists from the Trickest Resolvers [repository](https://github.com/trickest/resolvers).

## Installation
### Go
```
go install github.com/root4loot/publicresolvers/cmd/publicresolvers@latest
```

### Docker
```
git clone https://github.com/root4loot/publicresolvers.git && cd publicresolvers
docker build -t publicresolvers .
docker run -it publicresolvers -h
```

## Usage
```
Usage: publicresolvers [flag]

  -r,   --resolvers   fetch resolvers.txt             (resolver IP addresses)
  -t,   --trusted     fetch resolvers-trusted.txt     (trusted resolvers from organizations like Cloudflare, Google, etc.)
  -c,   --community   fetch resolvers-community.txt   (resolver IP addresses with community annotations)
  -h,   --help        display help
```


## Library

```
go get github.com/root4loot/publicresolvers
```

```go
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
```

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md)

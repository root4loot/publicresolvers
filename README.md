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
  -p,   --with-port   include port 53 in resolver IP addresses
  -h,   --help        display help
```

## Example
```
publicresolvers --resolvers

204.74.109.172
204.69.234.244
204.106.240.53
202.65.192.146
202.65.124.32
202.63.241.68
202.129.206.237
...
```

```
publicresolvers --resolvers --with-port

204.74.109.172:53
204.69.234.244:53
204.106.240.53:53
202.65.192.146:53
202.65.124.32:53
202.63.241.68:53
202.129.206.237:53
...
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

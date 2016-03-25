package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/skratchdot/open-golang/open"
)

const (
	google        = "https://google.com/?q=%s"
	duckduckGo    = "https://duckduckgo.com/?q=%s"
	bing          = "https://bing.com/?q=%s"
	commandAndArg = 2
)

func main() {

	//Setup flags
	query := flag.String("q", "", "Query string, wrapped in quotes if multi word")
	provider := flag.String("p", "google", "The search engine to perform the search")
	flag.Parse()

	// if no flag is provided we can still the next argument on command line as
	// query, then can prevent index bounds exception by not settings it as default
	if *query == "" && len(os.Args) >= commandAndArg {
		query = &os.Args[1]
	}

	//Execute matching the provider name
	if *provider != "" {
		providerMatched := matchProvider(*provider)
		open.Run(fmt.Sprintf(providerMatched, *query))
	} else if *query != "" {
		open.Run(fmt.Sprintf(google, *query))
	}

	fmt.Print("<< Nox >>")
}

func matchProvider(provider string) string {
	if strings.Contains(provider, "duck") {
		return duckduckGo
	} else if strings.Contains(provider, "bing") {
		return bing
	}
	return google
}

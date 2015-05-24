package main

import (
	"flag"
	"fmt"
	"os"
)

type config struct {
	Domain        string
	TargetBaseURL string
	Listen        string
}

func getConfig() *config {
	defaultListen := fmt.Sprintf(":%s", os.Getenv("POST"))
	if defaultListen == ":" {
		defaultListen = ":80"
	}

	var (
		domain        = flag.String("domain", os.Getenv("DOMAIN"), "Domain to filter for (ENV[DOMAIN])")
		targetBaseURL = flag.String("target", os.Getenv("TARGET"), "BaseURL to fetch content from (ENV[TARGET])")
		listen        = flag.String("listen", defaultListen, "Address to listen on")
	)

	flag.Parse()

	return &config{
		Domain:        *domain,
		TargetBaseURL: *targetBaseURL,
		Listen:        *listen,
	}
}

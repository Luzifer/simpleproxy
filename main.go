package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/elazarl/goproxy"
)

func main() {
	cfg := getConfig()

	proxy := goproxy.NewProxyHttpServer()
	proxy.OnRequest().HandleConnect(goproxy.AlwaysReject)

	// We are not really a proxy but act as a HTTP(s) server who delivers remote pages
	proxy.NonproxyHandler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		proxy.ServeHTTP(w, req)
	})

	http.ListenAndServe(cfg.Listen, shieldDomain(cfg, proxy))
}

func shieldDomain(cfg *config, handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// If the domain is set restrict access to that one domain
		// This is mainly to prevent open proxying
		if cfg.Domain != "" && req.Host != cfg.Domain {
			http.Error(w, fmt.Sprintf("This is not a public proxy. Access only for '%s', not for '%s'.", cfg.Domain, req.Host), http.StatusForbidden)
			return
		}

		target, err := url.Parse(cfg.TargetBaseURL)
		if err != nil {
			http.Error(w, "Unable to parse TargetURL", http.StatusInternalServerError)
		}

		target.Path = strings.Join([]string{target.Path, req.URL.Path}, "/")
		for strings.Contains(target.Path, "//") {
			target.Path = strings.Replace(target.Path, "//", "/", -1)
		}

		fmt.Printf("Target: %s\n", target.String())

		r, _ := http.NewRequest(req.Method, target.String(), req.Body)

		handler.ServeHTTP(w, r)
	})
}

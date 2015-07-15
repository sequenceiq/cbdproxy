package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/benschw/dns-clb-go/clb"
)

func dnsLb() (c clb.LoadBalancer) {

	dnsHost := os.Getenv("DNS_HOST")
	if dnsHost != "" {
		dnsPort := os.Getenv("DNS_PORT")
		if dnsPort == "" {
			dnsPort = "53"
		}

		c = clb.NewClb(dnsHost, dnsPort, clb.RoundRobin)
	} else {
		c = clb.NewDefaultClb(clb.RoundRobin)
	}
	return c
}

func getServiceUrl(s string) *url.URL {

	addr, err := dnsLb().GetAddress(s + ".service.consul")
	if err != nil {
		fmt.Fprintf(os.Stderr, "[WARNING] service: %s isnt registered yet ...", s)
		return nil
	}

	url, err := url.Parse("http://" + addr.String())
	if err != nil {
		fmt.Fprintf(os.Stderr, "[WARNING] cloudn't parse address: %s", addr.String())
		return nil
	}
	return url
}

func port() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	if port[0] != ':' {
		port = ":" + port
	}
	fmt.Println("[INFO] listening on:", port)
	return port
}

type ServiceProxyHandler struct {
	ServiceName  string
	Patt         string
	ServiceURL   *url.URL
	ProxyHandler http.Handler
}

func (h *ServiceProxyHandler) refreshURL() {
	u := getServiceUrl(h.ServiceName)
	if u != h.ServiceURL {
		h.ServiceURL = u
		p := httputil.NewSingleHostReverseProxy(u)
		h.ProxyHandler = http.StripPrefix(h.Patt, p)
	}
}

func (h *ServiceProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if os.Getenv("DEBUG") != "" {
		fmt.Printf("[DEBUG] PATT: %-10s REQ-URL:%s -> %s \n", h.Patt, r.URL.String(), h.ServiceURL.String())
	}

	h.refreshURL()

	if h.ServiceURL != nil {
		r.Header.Set("x-proxdnsLb()y-prefix", h.Patt)
		h.ProxyHandler.ServeHTTP(w, r)
	} else {
		fmt.Fprintf(os.Stderr, "[WARNING] ServiceURL is unknown ...\n")
	}
}

func NewServiceProxyHandler(service string, patt string) *ServiceProxyHandler {
	if patt == "" {
		patt = fmt.Sprintf("/%s/", service)
	}
	fmt.Println("[INFO] new proxy", patt, "->", service)
	h := &ServiceProxyHandler{
		ServiceName: service,
		Patt:        patt,
	}

	return h

}

func main() {

	for _, s := range []struct {
		service string
		patt    string
	}{
		{"cloudbreak", "/cloudbreak/"},
		{"identity", "/identity/"},
		{"uluwatu", "/uluwatu/"},
		{"sultans", "/sultans/"},
	} {
		http.Handle(s.service, NewServiceProxyHandler(s.service, s.patt))
	}

	http.ListenAndServe(port(), nil)
}

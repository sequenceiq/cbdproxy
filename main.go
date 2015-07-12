package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/benschw/dns-clb-go/clb"

	"golang.org/x/oauth2"
)

func getGithubConfig() *oauth2.Config {

	return &oauth2.Config{
		ClientID:     "953d27b8258dea6dd658",
		ClientSecret: "83a3590b8f6611304e1ae8bb595778db518f90f9",
		Scopes:       []string{"user", "gist"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://github.com/login/oauth/authorize",
			TokenURL: "https://github.com/login/oauth/access_token",
		},
	}
}
func getCloudbreakConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     "uluwatu",
		ClientSecret: "cbsecret2015",
		Scopes:       []string{"openid"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "http://b2d:8888/oauth/authorize",
			TokenURL: "http://b2d:8888/oauth/token",
		},
	}
}

func getToken() {

	fmt.Println("login to uaa ...")

	//conf := getGithubConfig()
	conf := getCloudbreakConfig()
	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	fmt.Printf("Visit the URL for the auth dialog: %v", url)

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatal(err)
	}
	tok, err := conf.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("GH_TOKEN=", tok.AccessToken, "  refresh=", tok.RefreshToken)
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{"version": 0.1}`)
}

type Todo struct {
	Todos []string
}

func NewTodo() *Todo {
	return &Todo{[]string{"one", "two", "tre"}}
}

func (t *Todo) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		for i, todo := range t.Todos {
			fmt.Fprintln(w, " -", i, ":", todo)
		}
	case "PUT", "POST":
		b, _ := ioutil.ReadAll(r.Body)
		t.Todos = append(t.Todos, string(b))
		fmt.Fprintln(w, "OK")
	}
}

func getServiceUrl(s string) *url.URL {
	c := clb.NewClb("b2d", "53", clb.RoundRobin)
	addr, err := c.GetAddress(s + ".service.consul")
	if err != nil {
		panic(err)
	}

	fmt.Println("[DEBUG] dns:", addr.String())

	url, err := url.Parse("http://" + addr.String())
	if err != nil {
		panic(err)
	}
	return url

}

func main() {
	//getToken()

	http.HandleFunc("/info", infoHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("."))))
	http.Handle("/todos/", NewTodo())

	for _, serv := range []string{"cloudbreak", "identity"} {
		patt := fmt.Sprintf("/%s/", serv)
		url := getServiceUrl(serv)
		fmt.Println(patt, "->", url)
		http.Handle(patt, http.StripPrefix(patt, httputil.NewSingleHostReverseProxy(url)))
	}

	http.ListenAndServe(":8080", nil)
}

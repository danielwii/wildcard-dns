package main

import "fmt"
import "os"
import "strings"
import "text/template"
// import (
// 	"io"
// )

const tmpl_hosts = `
$ttl 38400
{{.domain}}.	IN	SOA	{{.url}}. admin.{{.domain}}. (
			1465982662
			10800
			3600
			604800
			38400 )
{{.domain}}.		IN	NS	{{.url}}.
*.{{.domain}}.		IN	A	{{.url}}
{{.domain}}.		IN	A	{{.url}}
`
const tmpl_named = `
options {
	allow-recursion { any; };
	allow-recursion-on { any; };
};

server 223.5.5.5 {
};

{{range .}}
zone "{{.domain}}" {
	type master;
	file "/etc/bind/{{.domain}}.hosts";
};
{{end}}
`

func main() {
	urls := os.Getenv("URLS")
	domains := os.Getenv("DOMAINS")

	dnsList := dnsDefines(urls, domains)
	fmt.Println("result is", dnsList)
	fmt.Println("load all templates...")
	// fmt.Println(tmpl_hosts)
	// fmt.Println(tmpl_named)

	{
		fmt.Println("------------------------------------")
		tmpl, err := template.New("tmpl_named").Parse(tmpl_named)
		check(err)

		file, err := os.Create("/etc/bind/named.conf")
		check(err)

		tmpl.Execute(os.Stdout, dnsList)
		err1 := tmpl.Execute(file, dnsList)
		check(err1)
		fmt.Println("Generate named.conf file...")
	}


	{
		fmt.Println("------------------------------------")
		tmpl, err := template.New("tmpl_hosts").Parse(tmpl_hosts)
		check(err)

		for _, dns := range dnsList {
			// fmt.Println(dns)
			fmt.Println(dns["domain"])
			f, err := os.Create("/etc/bind/" + dns["domain"] + ".hosts")
			check(err)
			tmpl.Execute(os.Stdout, dns)
			tmpl.Execute(f, dns)
			fmt.Println("Generate hosts file(s)...")
		}
	}

	fmt.Println("DONE >_______________,.<")
}

func dnsDefines(urls string, domains string) []map[string]string {
	// fmt.Println(strings.Split(urls, ","))
	_urls := strings.Split(urls, ",")
	_domains := strings.Split(domains, ",")
	_size := len(_urls)
	// fmt.Println(strings.Split(domains, ","))
	x := make([]map[string]string, _size)
	// x := [_size]map
	for i := 0; i < _size; i++ {
		// fmt.Println(_urls[i] + ", " + _domains[i])
		_url := strings.TrimFunc(_urls[i], isBlank)
		_domain := strings.TrimFunc(_domains[i], isBlank)
		fmt.Println("-------->", _url + ", " + _domain)
		_wildcard_dns := make(map[string]string)
		_wildcard_dns["url"] = _url
		_wildcard_dns["domain"] = _domain
		fmt.Println(_wildcard_dns)
		x[i] = _wildcard_dns
	}
	fmt.Println("========>", x)
	return x
}

func isBlank(r rune) bool {
	return r == ' '
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

package main

import (
	"fmt"
	"net/url"
)

func main() {
	s := "https://www.google.com.ua/?gfe_rd=cr&ei=-nusV6eeDNXDtAG72a3YDg"

	u1, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(u1)
	fmt.Println(u1.Scheme)
	fmt.Println(u1.Host)
	fmt.Println(u1.Path)
	fmt.Println(u1.RawQuery)
	m1, _ := url.ParseQuery(u1.RawQuery)
	fmt.Println(m1)
	fmt.Println(m1["ei"][0])

	k := "postgres://userFOO:pass234234234@host.com:5432/path?k=v#f"

	u2, err := url.Parse(k)
	if err != nil {
		panic(err)
	}

	fmt.Println(u2)
	fmt.Println(u2.Scheme)
	fmt.Println(u2.User)
	fmt.Println(u2.User.Username())
	p, _ := u2.User.Password()
	fmt.Println(p)
	fmt.Println(u2.Host)
	fmt.Println(u2.Path)
	fmt.Println(u2.RawQuery)

	m2, _ := url.ParseQuery(u2.RawQuery)
	fmt.Println(m2)
	fmt.Println(m2["k"][0])

}

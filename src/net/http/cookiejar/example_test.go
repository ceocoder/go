// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cookiejar_test

import (
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
)

func ExampleNew() {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{Name: "COOKIE", Value: "NOM NOM NOM"})
	}))
	defer ts.Close()

	// A nil value is valid and may be useful for testing but it is not
	// secure: it means that the HTTP server for foo.co.uk can set a cookie
	// for bar.co.uk.
	//
	// Actual implementation in package golang.org/x/net/publicsuffix
	// jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatal(err)
	}

	client := http.Client{
		Jar: jar,
	}

	req, err := http.NewRequest("GET", ts.URL, nil)
	if _, err = client.Do(req); err != nil {
		log.Fatal(err)
	}

	for _, cookie := range jar.Cookies(req.URL) {
		fmt.Printf("%s:%s", cookie.Name, cookie.Value)
	}

	// Output:
	// COOKIE:NOM NOM NOM
}

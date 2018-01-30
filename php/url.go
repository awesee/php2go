package php

import (
	"net/url"
	"strings"
)

//Parse a URL and return its components
func ParseUrl(rawurl string) (*url.URL, error) {

	return url.Parse(rawurl)
}

//Parses the string into variables
func ParseStr(query string) (url.Values, error) {

	return url.ParseQuery(query)
}

//URL-encode according to RFC 3986
func Rawurlencode(s string) string {

	return url.PathEscape(s)
}

//Decodes URL-encoded string
func Rawurldecode(s string) (string, error) {

	return url.PathUnescape(s)
}

//URL-encodes string
func Urlencode(s string) string {

	return strings.Replace(url.PathEscape(s), "%20", "+", -1)
}

//Decodes URL-encoded string
func Urldecode(s string) (string, error) {

	return url.PathUnescape(strings.Replace(s, "+", "%20", -1))
}

package php

import (
	"net/url"
	"strings"
)

//ParseURL - Parse a URL and return its components
func ParseURL(rawurl string) (*url.URL, error) {

	return url.Parse(rawurl)
}

//ParseStr - Parses the string into variables
func ParseStr(query string) (url.Values, error) {

	return url.ParseQuery(query)
}

//Rawurlencode - URL-encode according to RFC 3986
func Rawurlencode(s string) string {

	return url.PathEscape(s)
}

//Rawurldecode - Decodes URL-encoded string
func Rawurldecode(s string) (string, error) {

	return url.PathUnescape(s)
}

//Urlencode - URL-encodes string
func Urlencode(s string) string {

	return strings.Replace(url.PathEscape(s), "%20", "+", -1)
}

//Urldecode - Decodes URL-encoded string
func Urldecode(s string) (string, error) {

	return url.PathUnescape(strings.Replace(s, "+", "%20", -1))
}

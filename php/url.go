package php

import (
	"net/url"
)

//Parse a URL and return its components
func ParseUrl(rawurl string) (*url.URL, error) {

	return url.Parse(rawurl)
}

//Parses the string into variables
func ParseStr(query string) (url.Values, error) {

	return url.ParseQuery(query)
}

//URL-encodes string
func Urlencode(s string) string {

	return url.PathEscape(s)
}

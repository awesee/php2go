package php

import "net/url"

func ParseUrl(s string) (*url.URL, error) {

	return url.Parse(s)
}

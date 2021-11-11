package dstp

import (
	"fmt"
	"net"
	"net/url"
	"strings"
)

func getAddr(addr string) (string, error) {
	var pu string

	// Cleanup the scheme first
	//
	// [scheme:][//[userinfo@]host][/]path[?query][#fragment]
	for _, prefix := range []string{"https://", "http://"} {
		if strings.HasPrefix(addr, prefix) {
			addr = strings.ReplaceAll(addr, prefix, "")
		}
	}

	ip := net.ParseIP(addr)
	if ip == nil {
		// This case is only for URL's
		// add a scheme for conform go's url form
		addr = "https://" + addr
		parsedURL, err := url.ParseRequestURI(addr)
		if err != nil {
			u, err := url.ParseRequestURI(addr)
			if err != nil {
				host, _, err := net.SplitHostPort(addr)
				if err != nil {
					return "", fmt.Errorf("failed to split host and port")
				}

				uu, err := url.ParseRequestURI(host)
				if err != nil {
					return "", fmt.Errorf("failed to parse url: %v", err.Error())
				}
				pu = uu.Hostname()
			} else {
				pu = u.Host
			}
		} else {
			// If URI doesn't have any slash in it, the first part is considered as scheme
			if parsedURL.Hostname() == "" {
				pu = parsedURL.Scheme
			} else {
				pu = parsedURL.Hostname()
			}
		}
	} else {
		pu = addr
	}

	return pu, nil
}

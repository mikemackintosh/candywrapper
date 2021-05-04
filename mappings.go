package main

import (
	"fmt"
	"net"
	"regexp"
	"strings"
)

/*
  cw=v1 from[*.mikemackintosh] to[mikemackintosh.com]
  cw=(?=ver())
  from[(.*)]
  to[(.*)]
*/

func validateDNS(hostHeader string) bool {
	fmt.Printf("Validating [%s]\n", hostHeader)
	split := strings.Split(hostHeader, ".")
	domain := strings.Join(split[len(split)-2:], ".")
	fmt.Printf("-> Looking up [%s]\n ", domain)
	txtrecords, err := net.LookupTXT(domain)
	if err != nil {
		return false
	}

	fmt.Println("\tfound records:")
	for _, txt := range txtrecords {
		if strings.HasPrefix(txt, "cw=v1") && strings.HasSuffix(txt, hostHeader) {
			fmt.Println(parseRecord(txt))
		}
	}

	return true
}

// Record defines a new record.
type Record struct {
	Include string `cw:"include"`
	To      string `cw:"to"`
}

// parseRecord will parse the TXT record.
func parseRecord(txt string) *Record {
	record := &Record{}
	r := regexp.MustCompile(`cw=v1 (include:(?P<include>[a-zA-Z0-9_\-.]+)|includeAll:(?P<includeAll>[a-zA-Z0-9_\-.]+)) to(?P<to>[a-zA-Z0-9_\-.]+)`)
	fmt.Printf("%#v\n", r.FindStringSubmatch(txt))
	fmt.Printf("%#v\n", r.SubexpNames())
	return record
}

/*
When a request is received by CandyWrapper, it will look at the Host request header. It will then take this request header and match it against a destination domain. If `includePath` is set in the `cw` wrapper.

How to validate something can be forwarded to.
www.mikemackintosh.com . www.mikemackintosh.com -> mikemackintosh.com.
mikemackintosh.com verifies www.mikemackintosh.com as a referrer.
TXT cw.mikemackintosh.com would then be verified
-> cw=v1 allowed[*.mikemackintosh.com]

Does this prevent:
www.mikemackintosh from redirecting to malicious-domain.com?
Yes, the server would check: [www.mikemackintosh.com]

cw=v1 sub:www

Host header, get domain.tld.
lookup cw records, and validate that: Host subdomain match

## Configuring your DNS record:
Configuring the DNS record

If you'd like to use this service, please PayPal me $5 for a lifetime membership with your Pull Request number including the mapping of your domain.




*/

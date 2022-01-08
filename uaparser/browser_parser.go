package uaparser

import (
	"regexp"
	"strings"
)

const (
	AMAZON     = "Amazon"
	APPLE      = "Apple"
	ASUS       = "ASUS"
	BLACKBERRY = "BlackBerry"
	BROWSER    = "Browser"
	CHROME     = "Chrome"
	EDGE       = "Edge"
	FIREFOX    = "Firefox"
	GOOGLE     = "Google"
	HUAWEI     = "Huawei"
	LG         = "LG"
	MICROSOFT  = "Microsoft"
	MOTOROLA   = "Motorola"
	OPERA      = "Opera"
	SAMSUNG    = "Samsung"
	SONY       = "Sony"
	XIAOMI     = "Xiaomi"
	ZEBRA      = "Zebra"
	FACEBOOK   = "Facebook"
	IE         = "IE"
)

type BrowserParser struct {
	Expr []*regexp.Regexp

	VersionIdx    int
	VersionParser func([]byte) string

	Name       string
	NameIndex  int
	NameParser func([]byte) string
}

func (r *BrowserParser) Match(b []byte) bool {
	for _, expr := range r.Expr {
		if expr.Match(b) {
			return true
		}
	}
	return false
}

func (r *BrowserParser) Parse(b []byte) *Browser {
	broswer := &Browser{}

	for _, expr := range r.Expr {
		if expr.Match(b) {
			match := expr.FindSubmatch(b)
			broswer.Name = r.ParseName(match)
			broswer.Version = r.ParseVersion(match)
			broswer.Major = r.ParseMajor(broswer.Version)
			//fmt.Printf("expr match here is : %s, name is %s\n", expr.String(), broswer.Name)
			break
		}
	}

	return broswer
}

func (r *BrowserParser) ParseName(match [][]byte) string {

	if r.Name != "" {
		return r.Name
	}

	if r.NameIndex < len(match) {
		if r.NameParser != nil {
			return r.NameParser(match[r.NameIndex])
		}
		return string(match[r.NameIndex])
	}

	return "undefined"
}

func (r *BrowserParser) ParseVersion(match [][]byte) string {
	if r.VersionIdx < len(match) {
		if r.VersionParser != nil {
			return r.VersionParser(match[r.VersionIdx])
		}
		if string(match[r.VersionIdx]) == "" {
			return "undefined"
		}
		return string(match[r.VersionIdx])
	}

	return "undefined"
}

func (r *BrowserParser) ParseMajor(version string) string {
	if version != "" && version != "undefined" {
		reg := regexp.MustCompile(`[^\d\.]`)
		return strings.Split(reg.ReplaceAllString(version, ""), ".")[0]
	}
	return "undefined"
}

// old version to x.x.x
var safariMap map[string]string = map[string]string{
	"1.0":   "/8",
	"1.2":   "/1",
	"1.3":   "/3",
	"2.0":   "/412",
	"2.0.2": "/416",
	"2.0.3": "/417",
	"2.0.4": "/419",
	"?":     "/",
}

func SafariMap(b []byte) string {
	version := string(b)

	for fix, old := range safariMap {
		if strings.Index(version, old) >= 0 {
			return fix
		}
	}
	return version
}

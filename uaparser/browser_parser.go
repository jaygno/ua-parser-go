package uaparser

import (
	"regexp"
	"strings"
)

var regMajor = regexp.MustCompile(`[^\d\.]`)

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

	UNDEFINED = "undefined"
)

// BrowserParser
type BrowserParser struct {
	Expr []*regexp.Regexp

	VersionIdx    int
	VersionParser func([]byte) string

	Name       string
	NameIndex  int
	NameParser func([]byte) string
}

// Match
func (r *BrowserParser) Match(b []byte) bool {
	for _, expr := range r.Expr {
		if expr.Match(b) {
			return true
		}
	}
	return false
}

// Parse
func (r *BrowserParser) Parse(b []byte) *Browser {
	broswer := &Browser{}

	for _, expr := range r.Expr {
		if expr.Match(b) {
			match := expr.FindSubmatch(b)
			broswer.Name = r.ParseName(match)
			broswer.Version = r.ParseVersion(match)
			broswer.Major = r.ParseMajor(broswer.Version)
			//fmt.Println(expr.String(), broswer.Version)
			//fmt.Printf("expr match here is : %s, name is %s\n", expr.String(), broswer.Name)
			break
		}
	}

	return broswer
}

// ParseName
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

	return UNDEFINED
}

// ParseVersion
func (r *BrowserParser) ParseVersion(match [][]byte) string {
	if r.VersionIdx < len(match) {
		if r.VersionParser != nil {
			return r.VersionParser(match[r.VersionIdx])
		}
		if string(match[r.VersionIdx]) == "" {
			return UNDEFINED
		}
		return string(match[r.VersionIdx])
	}

	return UNDEFINED
}

// ParseMajor
func (r *BrowserParser) ParseMajor(version string) string {
	if version != "" && version != UNDEFINED {
		return strings.Split(regMajor.ReplaceAllString(version, ""), ".")[0]
	}
	return UNDEFINED
}

// old version to x.x.x
var safariMap = map[string]string{
	"1.0":   "/8",
	"1.2":   "/1",
	"1.3":   "/3",
	"2.0":   "/412",
	"2.0.2": "/416",
	"2.0.3": "/417",
	"2.0.4": "/419",
}

// oldSafariParser
func oldSafariParser(b []byte) string {
	version := string(b)

	for fix, old := range safariMap {
		if strings.Index(version, old) >= 0 {
			return fix
		}
	}
	return version
}

package generator

import (
	"strings"

	"github.com/gertd/go-pluralize"
	"github.com/stoewer/go-strcase"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Field string

// https://github.com/morix1500/lint/blob/6aaf7c34af0f4c36a57e0c429bace4d706d8e931/lint.go#L740-L782
var commonInitialisms = map[string]bool{
	"ACL":   true,
	"API":   true,
	"ASCII": true,
	"CPU":   true,
	"CSS":   true,
	"DNS":   true,
	"EOF":   true,
	"GUID":  true,
	"HTML":  true,
	"HTTP":  true,
	"HTTPS": true,
	"ID":    true,
	"IP":    true,
	"JSON":  true,
	"LHS":   true,
	"QPS":   true,
	"RAM":   true,
	"RHS":   true,
	"RPC":   true,
	"SLA":   true,
	"SMTP":  true,
	"SQL":   true,
	"SSH":   true,
	"TCP":   true,
	"TLS":   true,
	"TTL":   true,
	"UDP":   true,
	"UI":    true,
	"UID":   true,
	"UUID":  true,
	"URI":   true,
	"URL":   true,
	"UTF8":  true,
	"VM":    true,
	"XML":   true,
	"XMPP":  true,
	"XSRF":  true,
	"XSS":   true,
}

func (f Field) String() string {
	return string(f)
}

func (f Field) ToUpperCamel() Field {
	if len(f) == 0 {
		return ""
	}

	snake := strcase.SnakeCase(f.String())
	words := strings.Split(snake, "_")

	results := make([]string, len(words))
	for _, word := range words {
		upperWord := strings.ToUpper(word)
		caser := cases.Title(language.English)

		if commonInitialisms[upperWord] {
			results = append(results, upperWord)
		} else {
			results = append(results, caser.String(word))
		}
	}

	return Field(strings.Join(results, ""))
}

func (f Field) ToSingular() Field {
	plc := pluralize.NewClient()

	return Field(plc.Singular(f.String()))
}

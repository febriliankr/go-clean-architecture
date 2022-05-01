package utils

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func Commafy(nominal int64) string {
	p := message.NewPrinter(language.Indonesian)
	return p.Sprintf("%d", nominal)
}

func Rupiahify(nominal int64) string {
	p := message.NewPrinter(language.Indonesian)
	return p.Sprintf("Rp%d", nominal)
}

func Rupiahify00(nominal int64) string {
	p := message.NewPrinter(language.Indonesian)
	return p.Sprintf("Rp%d,00", nominal)
}


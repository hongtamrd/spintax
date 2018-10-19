package spintax

import (
	"math/rand"
	"strings"
)

type (
	Exp []Spintax
	Alt []Spintax
	Str string

	Spintax interface {
		Spin() string
		Count() int
		String() string
	}
)

func Parse(exp string) Spintax {
	e, _ := parseExp(exp, 0)
	return e
}

func parseExp(e string, i int) (Spintax, int) {
	var r Exp
	s := i
loop:
	for i < len(e) {
		c := e[i]
		switch c {
		case '{':
			if s != i {
				r = append(r, Str(e[s:i]))
			}
			var alt Spintax
			alt, i = parseAlt(e, i)
			i++
			s = i
			if alt != nil {
				r = append(r, alt)
			}
		case '|', '}':
			break loop
		default:
			i++
		}
	}
	if s != i {
		r = append(r, Str(e[s:i]))
	}
	if r == nil {
		return Str(""), i
	}
	if len(r) == 1 {
		return r[0], i
	}
	return r, i
}

func parseAlt(e string, i int) (Spintax, int) {
	var r Alt
	var exp Spintax
	for i < len(e) {
		if e[i] == '}' {
			break
		}
		if e[i] == '|' || e[i] == '{' {
			i++
		}
		exp, i = parseExp(e, i)
		r = append(r, exp)
	}
	if len(r) == 1 {
		return r[0], i
	}
	return r, i
}

func (e Exp) Spin() string {
	var b strings.Builder
	for _, e := range e {
		b.WriteString(e.Spin())
	}
	return b.String()
}

func (a Alt) Spin() string {
	e := a[rand.Intn(len(a))]
	return e.Spin()
}

func (s Str) Spin() string { return string(s) }

func (e Exp) Count() int {
	s := 1
	for _, e := range e {
		s *= e.Count()
	}
	return s
}

func (a Alt) Count() int {
	s := 0
	for _, e := range a {
		s += e.Count()
	}
	return s
}

func (s Str) Count() int { return 1 }

func (e Exp) String() string {
	var b strings.Builder
	for _, e := range e {
		b.WriteString(e.String())
	}
	return b.String()
}

func (a Alt) String() string {
	var b strings.Builder
	b.WriteString("{")
	for i, e := range a {
		if i != 0 {
			b.WriteString("|")
		}
		b.WriteString(e.String())
	}
	b.WriteString("}")
	return b.String()
}

func (s Str) String() string { return string(s) }
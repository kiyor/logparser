/* -.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.

* File Name : logparser.go

* Purpose :

* Creation Date : 06-02-2015

* Last Modified : Wed 24 Jun 2015 01:17:43 AM UTC

* Created By : Kiyor

_._._._._._._._._._._._._._._._._._._._._.*/

package logparser

import (
	"errors"
	"strings"
)

type Parser struct {
	pi *[]*pair
}

type pair struct {
	L rune
	R rune
}

func NewParser(rs ...rune) (*Parser, error) {
	if len(rs)%2 != 0 {
		return nil, errors.New("input rune shoud be pair")
	}
	var p Parser
	var pi []*pair
	p.pi = &pi
	var lpair *pair = new(pair)
	for k, v := range rs {
		if k%2 == 0 {
			lpair.L = v
		}
		if k%2 == 1 {
			lpair.R = v
			*p.pi = append(*p.pi, lpair)
			lpair = new(pair)
		}
	}
	return &p, nil
}

func (p *Parser) lin(r rune) (bool, *rune) {
	for _, p := range *p.pi {
		if r == p.L {
			return true, &p.R
		}
	}
	return false, nil
}
func (p *Parser) rin(r rune) bool {
	for _, p := range *p.pi {
		if r == p.R {
			return true
		}
	}
	return false
}

func (p *Parser) Split(line string) []string {
	var out []string
	var word string
	var pair rune = 0
	var needRmatch bool
	for _, c := range line {
		// if right match
		if pair == c {
			out = aappend(out, word)
			word = ""
			pair = 0
			if p.rin(c) && needRmatch {
				needRmatch = false
				continue
			}
			word += string(c)
		} else {
			// if space match and just finish word
			if c == ' ' {
				if needRmatch {
					word += string(c)
					continue
				}
				out = aappend(out, word)
				word = ""
			} else if b, r := p.lin(c); b && !needRmatch { // if left match
				out = aappend(out, word)
				pair = *r
				needRmatch = true
			} else {
				word += string(c)
			}
		}
	}
	out = aappend(out, word)

	return out
}

func aappend(ss []string, s string) []string {
	s = strings.Trim(s, " ")
	if len(s) > 0 {
		ss = append(ss, s)
	}
	return ss
}

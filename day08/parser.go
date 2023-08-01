package main

import "log"

type Parser struct {
	str       string
	pos       int
	byteCount int
}

func (p *Parser) char() byte {
	return p.str[p.pos]
}

func (p *Parser) countBytes() int {
	if len(p.str) <= 1 {
		return 0
	}
	if p.char() == '"' {
		p.pos += 1
		p.inQuote()
	}
	return p.byteCount
}

func (p *Parser) inQuote() {
	for {
		if p.char() == '\\' {
			p.pos += 1
			p.inBackslash()
			continue
		}
		if p.char() == '"' {
			p.pos += 1
			return
		}
		p.pos += 1
		p.byteCount += 1
	}
}

func (p *Parser) inBackslash() {
	if p.char() == '"' || p.char() == '\\' {
		p.pos += 1
		p.byteCount += 1
	} else if p.char() == 'x' {
		p.pos += 3
		p.byteCount += 1
	} else {
		log.Fatal("Parse error")
	}
}

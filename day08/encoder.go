package main

type Encoder struct {
	str       string
	bytecount int
}

func (p *Encoder) encode() {
	p.bytecount = 2
	for _, v := range p.str {
		if v == '"' {
			p.bytecount += 2
		} else if v == '\\' {
			p.bytecount += 2
		} else {
			p.bytecount += 1
		}
	}
}

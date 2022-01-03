package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type snailfish struct {
	h      int        // height
	m      int        // maximum value in this subtree
	v      int        // value (valid only if lc == rc == nil)
	lc, rc *snailfish // left / right child
	p      *snailfish // parent
}

func (f *snailfish) String() string {
	if f.lc == nil && f.rc == nil {
		return fmt.Sprintf("%d", f.v)
	}

	var sb strings.Builder
	sb.WriteByte('[')
	sb.WriteString(f.lc.String())
	sb.WriteByte(',')
	sb.WriteString(f.rc.String())
	sb.WriteByte(']')
	return sb.String()
}

func read(s string) (*snailfish, string) {
	f := &snailfish{}
	if s[1] == '[' {
		f.lc, s = read(s[1:])
		f.lc.p = f
		f.h = f.lc.h + 1
		f.m = f.lc.m
	} else {
		v := int(s[1]) - int('0')
		f.lc = &snailfish{
			m: v,
			v: v,
			p: f,
		}
		f.h = 1
		f.m = v
		s = s[2:]
	}
	if s[1] == '[' {
		f.rc, s = read(s[1:])
		f.rc.p = f
		if f.rc.h+1 > f.h {
			f.h = f.rc.h + 1
		}
		if f.rc.m > f.m {
			f.m = f.rc.m
		}
	} else {
		v := int(s[1]) - int('0')
		f.rc = &snailfish{
			m: v,
			v: v,
			p: f,
		}
		if v > f.m {
			f.m = v
		}
		s = s[2:]
	}
	return f, s[1:]
}

func add(lf, rf *snailfish) *snailfish {
	f := &snailfish{
		lc: lf,
		rc: rf,
	}
	lf.p = f
	rf.p = f
	f.h = lf.h + 1
	if rf.h+1 > f.h {
		f.h = rf.h + 1
	}
	f.m = lf.m
	if rf.m > f.m {
		f.m = rf.m
	}
	return f
}

func explode(f *snailfish) *snailfish {
	h := 5
	if f == nil || f.h != h {
		return nil
	}

	for f.h == h && h > 1 {
		if f.lc != nil && f.lc.h == h-1 {
			f = f.lc
		} else if f.rc != nil && f.rc.h == h-1 {
			f = f.rc
		}
		h--
	}
	return f
}

func split(f *snailfish) *snailfish {
	if f == nil || f.m < 10 {
		return nil
	}

	for f.m >= 10 && f.lc != nil && f.rc != nil {
		if f.lc.m >= 10 {
			f = f.lc
		} else {
			f = f.rc
		}
	}
	return f
}

func left(f *snailfish) *snailfish {
	for f.p != nil && f == f.p.lc {
		f = f.p
	}
	if f.p != nil && f == f.p.rc {
		f = f.p.lc
		for f.rc != nil {
			f = f.rc
		}
		return f
	}
	return nil
}

func right(f *snailfish) *snailfish {
	for f.p != nil && f == f.p.rc {
		f = f.p
	}
	if f.p != nil && f == f.p.lc {
		f = f.p.rc
		for f.lc != nil {
			f = f.lc
		}
		return f
	}
	return nil
}

func update(f *snailfish) {
	for f.p != nil {
		f = f.p
		f.h = f.lc.h + 1
		if f.rc.h+1 > f.h {
			f.h = f.rc.h + 1
		}
		f.m = f.lc.m
		if f.rc.m > f.m {
			f.m = f.rc.m
		}
	}
}

func replace(f *snailfish, g *snailfish) {
	g.p = f.p
	if f.p == nil {
		return
	}

	if f == f.p.lc {
		f.p.lc = g
	} else if f == f.p.rc {
		f.p.rc = g
	}
	update(g)
}

func pair(n int) *snailfish {
	f := &snailfish{
		h: 1,
		m: n/2 + n%2,
		lc: &snailfish{
			m: n / 2,
			v: n / 2,
		},
		rc: &snailfish{
			m: n/2 + n%2,
			v: n/2 + n%2,
		},
	}
	f.lc.p = f
	f.rc.p = f
	return f
}

func reduce(f *snailfish) bool {
	if g := explode(f); g != nil {
		l := left(g)
		r := right(g)
		if l != nil {
			l.v += g.lc.v
			l.m = l.v
			update(l)
		}
		if r != nil {
			r.v += g.rc.v
			r.m = r.v
			update(r)
		}
		replace(g, &snailfish{})
		return true
	} else if g := split(f); g != nil {
		replace(g, pair(g.v))
		return true
	}
	return false
}

func magnitude(f *snailfish) int {
	if f.lc == nil && f.rc == nil {
		return f.v
	}

	return 3*magnitude(f.lc) + 2*magnitude(f.rc)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	l := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l = append(l, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	max := 0
	for i := 0; i < len(l); i++ {
		for j := 0; j < len(l); j++ {
			if i == j {
				continue
			}

			f, _ := read(l[i])
			g, _ := read(l[j])
			h := add(f, g)
			for reduce(h) {
			}

			m := magnitude(h)
			if m > max {
				max = m
			}
		}
	}

	fmt.Println(max)
}

package seqx

import "strings"

//https://stackoverflow.com/questions/38923376/return-a-new-string-that-sorts-between-two-given-strings

const beforeA = 'a' - 1
const afterZ = 'z' + 1

// Middle return a string that sort between prev & next
func Middle(prev, next string) string {
	var p, n byte
	pos := 0
	var str strings.Builder

	// copy identical part to str
	for ; p == n; pos++ {
		p = beforeA
		if pos < len(prev) {
			p = prev[pos]
		}

		n = afterZ
		if pos < len(next) {
			n = next[pos]
		}

		if p == n {
			str.WriteByte(p)
		}
	}

	if p == beforeA { // end of left string
		for n == 'a' { // handle a's
			n = afterZ
			if pos < len(next) {
				n = next[pos]
				pos++
			}

			str.WriteByte('a')
		}
		if n == 'b' {
			str.WriteByte('a')
			n = afterZ
		}
	} else if p+1 == n {
		str.WriteByte(p)

		n = afterZ
		for {
			p = beforeA
			if pos < len(prev) {
				p = prev[pos]
				pos++
			}

			if p != 'z' {
				break
			}

			str.WriteByte('z')
		}
	}

	str.WriteByte(n - (n-p)/2) //append middle character

	return str.String()
}

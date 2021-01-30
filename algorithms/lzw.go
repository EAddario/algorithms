package algorithms

import "../stdin"

// LZW ...
type LZW struct {
	R int // number of input chars
	L int // number of codewords = 2^W
	W int // codeword width
}

// NewLZW ...
func NewLZW() *LZW {
	return &LZW{256, 4096, 12}
}

// Compress ...
func (l *LZW) Compress() {
	input := stdin.BinaryStdin.ReadString()
	st := NewTST()

	for i := 0; i < l.R; i++ {
		st.Put(string(rune(i)), i)
	}

	code := l.R + 1 // R is codeword for EOF

	for len(input) > 0 {
		s := st.LongPrefixOf(input)                      // Find max prefix match s.
		_ = stdin.BinaryStdout.WriteBits(st.Get(s).(int), l.W) // Print s's encoding.
		t := len(s)

		if t < len(input) && code < l.L {
			st.Put(input[0:t+1], code)
			code++
		}

		input = input[t:]
	}

	_ = stdin.BinaryStdout.WriteBits(l.R, l.W)
	stdin.BinaryStdout.Close()
}

// Expand ...
func (l *LZW) Expand() {
	st := make([]string, l.L)
	var i int

	for i = 0; i < l.R; i++ {
		st[i] = string(rune(i))
	}

	st[i] = " "
	i++
	codeword := stdin.BinaryStdin.ReadIntR(l.W)
	val := st[codeword]

	for {
		_ = stdin.BinaryStdout.WriteStr(val)
		codeword = stdin.BinaryStdin.ReadIntR(l.W)

		if codeword == l.R {
			break
		}

		s := st[codeword]

		if i == codeword {
			s = val + string(val[0])
		}

		if i < l.L {
			st[i] = val + string(s[0])
			i++
		}

		val = s
	}

	stdin.BinaryStdout.Close()
}

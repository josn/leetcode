package algs0006

import (
    "bytes"
)

func convert(s string, numRows int) string {
   if numRows == 1 || len(s) <= numRows {
		return s
	}

	res := bytes.Buffer{}

	p := numRows*2 - 2

	// first line
	for i := 0; i < len(s); i += p {
		res.WriteByte(s[i])
	}

	// lines between first line and last line
	for r := 1; r <= numRows-2; r++ {
		// 添加r行的第一个字符
		res.WriteByte(s[r])

		for k := p; k-r < len(s); k += p {
			res.WriteByte(s[k-r])
			if k+r < len(s) {
				res.WriteByte(s[k+r])
			}
		}
	}

	// last line
	for i := numRows - 1; i < len(s); i += p {
		res.WriteByte(s[i])
	}

	return res.String()
}

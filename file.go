package totext

import (
	"bufio"
	"io"
)

// ReadLine is function for iterate lines on bufio.Reader.
// The example value for limitBytes is 1000000.
func ReadLine(rd *bufio.Reader, limitBytes int) (string, bool, error) {
	iseof := false
	buf := make([]byte, 0, limitBytes)

	for {
		l, p, e := rd.ReadLine()
		if e != nil {
			if e == io.EOF {
				iseof = true
				break
			} else {
				return "", false, e
			}
		}

		buf = append(buf, l...)

		if !p {
			break
		}
	}

	return string(buf), iseof, nil
}

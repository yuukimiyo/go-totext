package totext

import (
	"bufio"
	"io"
)

// ReadLineOld is old function for iterate lines on bufio.Reader.
// The example value for limitBytes is 1000000.
func ReadLineOld(rd *bufio.Reader, limitBytes int) (string, bool, error) {
	iseof := false
	buf := make([]byte, 0, limitBytes)

	for {
		l, p, e := rd.ReadLine()
		buf = append(buf, l...)

		if e != nil {
			if e == io.EOF {
				iseof = true
				break
			} else {
				return string(buf), false, e
			}
		}

		if !p {
			break
		}
	}

	return string(buf), iseof, nil
}

// ReadLine is new function for iterate lines on bufio.Reader.
// lineBuffer is byte slice, you can use as make([]byte, 0, 1024*1024) or more simply []byte.
// example:
// var rd = bufio.NewReaderSize(fp, bufferSize)
// for {
//     line, err := readLine(rd, make([]byte, 0, 1024*1024))
//     if err != nil {
//         if err == io.EOF {
//             break
//         }
//         panic(err)
//     }
//     lines = append(lines, line)
// }
func ReadLine(reader *bufio.Reader, lineBuffer []byte) (string, error) {
	for {
		l, hasNext, e := reader.ReadLine()
		lineBuffer = append(lineBuffer, l...)

		if !hasNext || e != nil {
			return string(lineBuffer), e
		}
	}
}

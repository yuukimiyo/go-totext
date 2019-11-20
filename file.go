package totext

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// ReadLine is function to iterate lines on bufio.Reader.
// lineBuffer should be byte slice to buffer data, you can make it as make([]byte, 0, 1024*1024) or more simply []byte.
// 3rd option number of 'make()' will be effect to read speed,
// It's depends on average size of length of each lines, I think.
// In many cases it's enough to 1024*1024.
//
// usage:
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

// ReadLineNormal is old function to iterate lines on bufio.Reader.
// *** this is old function. ***
// The example value for limitBytes is 1000000.
func ReadLineNormal(rd *bufio.Reader, limitBytes int) (string, bool, error) {
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

// WriteLines is function to write string array.
func WriteLines(filename string, lines []string, linesep string, modeflag string, permission os.FileMode) error {
	mode := os.O_WRONLY | os.O_CREATE
	if modeflag == "a" {
		mode = os.O_WRONLY | os.O_APPEND
	}

	f, err := os.OpenFile(filename, mode, permission)
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)

	if linesep == "" {
		// 改行文字列が空白だった場合
		for _, line := range lines {
			fmt.Print(w, line)
		}
	} else {
		for _, line := range lines {
			fmt.Println(w, line)
		}
	}

	return nil
}

// MakeDir is function to make dir if not exists.
func MakeDir(path string) error {
	err := os.Mkdir(path, 0755)
	if err == nil || os.IsExist(err) {
		return nil
	}

	return err
}

// Dirs is function to get dir list from inputted path.
func Dirs(dataRoot string) ([]string, error) {
	var dataDirs []string

	files, err := ioutil.ReadDir(dataRoot)
	if err != nil {
		return dataDirs, err
	}

	for _, file := range files {
		if file.IsDir() {
			dataDirs = append(dataDirs, filepath.Join(dataRoot, file.Name()))
		}
	}

	return dataDirs, nil
}

// Files is function to get file list from inputted path.
func Files(dataRoot string) ([]string, error) {
	var dataFiles []string

	files, err := ioutil.ReadDir(dataRoot)
	if err != nil {
		return dataFiles, err
	}

	for _, f := range files {
		fullpath := filepath.Join(dataRoot, f.Name())
		// _, err := os.Stat(fullpath)
		// if err != nil {
		// }
		dataFiles = append(dataFiles, fullpath)
	}

	return dataFiles, nil
}

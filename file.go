package totext

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// ReadLine is new function to iterate lines on bufio.Reader.
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

// ReadLineOld is old function to iterate lines on bufio.Reader.
// *** this is old function. ***
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

	for _, line := range lines {
		_, err := f.WriteString(line + linesep)
		if err != nil {
			return err
		}
	}

	return nil
}

//MakeDir is function to make dir.
// Make dir if not exists.
func MakeDir(path string) error {
	err := os.Mkdir(path, 0755)
	if err == nil || os.IsExist(err) {
		return nil
	}
	return err
}

// Dirs is function to get dir list from dirpath.
func Dirs(dataRoot string) []string {
	glog.V(3).Infof("datadir: " + dataRoot)

	var dataDirs []string

	files, err := ioutil.ReadDir(dataRoot)
	if err != nil {
		glog.Error(err)
	}

	for _, file := range files {
		if file.IsDir() {
			dataDirs = append(dataDirs, filepath.Join(dataRoot, file.Name()))
		}
	}

	return dataDirs
}

// Files is function to get file list from file dir.
func Files(dataRoot string) []string {
	var dataFiles []string

	files, err := ioutil.ReadDir(dataRoot)
	if err != nil {
		glog.Error(err)
	}

	for _, f := range files {
		fullpath := filepath.Join(dataRoot, f.Name())
		// _, err := os.Stat(fullpath)
		// if err != nil {
		// }
		dataFiles = append(dataFiles, fullpath)
	}

	return dataFiles
}

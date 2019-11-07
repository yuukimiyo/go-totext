package totext

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"io"
)

// Deflate is function for deflate text.
func Deflate(text string) (string, error) {
	zlibBuffer, err := compressStringToZlib(text)
	if err != nil {
		panic(err)
	}

	base64String := convertBytesToBase64String(zlibBuffer.Bytes())

	return base64String, nil
}

func compressStringToZlib(text string) (*bytes.Buffer, error) {
	textBuffer := bytes.NewBufferString(text)
	zlibBuffer := new(bytes.Buffer)
	zlibWriter := zlib.NewWriter(zlibBuffer)

	defer zlibWriter.Close()

	if _, err := io.Copy(zlibWriter, textBuffer); err != nil {
		return nil, err
	}

	return zlibBuffer, nil
}

func convertBytesToBase64String(b []byte) string {
	base64String := base64.StdEncoding.EncodeToString(b)
	return base64String
}

// Inflate is function for inflate base64text.
func Inflate(base64String string) (string, error) {

	bytes, err := convertBase64StringToBytes(base64String)
	if err != nil {
		return "", err
	}

	inflated, err := decompressZlibToString(bytes)
	if err != nil {
		return "", err
	}

	return inflated, nil
}

func convertBase64StringToBytes(base64String string) ([]byte, error) {
	bytes, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func decompressZlibToString(zlibBytes []byte) (string, error) {
	bytesReader := bytes.NewReader(zlibBytes)

	zlibReader, err := zlib.NewReader(bytesReader)
	if err != nil {
		return "", err
	}

	zlibBuffer := new(bytes.Buffer)
	err := zlibBuffer.ReadFrom(zlibReader)
	if err != nil {
		return "", err
	}

	return zlibBuffer.String(), nil
}

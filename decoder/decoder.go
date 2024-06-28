package decoder

import (
	"bytes"
	"io"
	"regexp"
	"strings"

	"encoding/base64"
	"mime/quotedprintable"

	"github.com/paulrosania/go-charset/charset"
)

func UTF8(cs string, data []byte) ([]byte, error) {
	if strings.ToUpper(cs) == "UTF-8" {
		return data, nil
	}

	r, err := charset.NewReader(cs, bytes.NewReader(data))
	if err != nil {
		return []byte{}, err
	}

	return io.ReadAll(r)

}

func Parse(bstr []byte) ([]byte, error) {
	var err error
	strs := regexp.MustCompile("^=\\?(.*?)\\?(.*?)\\?(.*)\\?=$").FindAllStringSubmatch(string(bstr), -1)

	if len(strs) > 0 && len(strs[0]) == 4 {
		c := strs[0][1]
		e := strs[0][2]
		dstr := strs[0][3]

		bstr, err = Decode(e, []byte(dstr))
		if err != nil {
			return bstr, err
		}

		return UTF8(c, bstr)
	}
	return bstr, err

}

func Decode(e string, bstr []byte) ([]byte, error) {
	var err error
	switch strings.ToUpper(e) {
	case "Q":
		bstr, err = io.ReadAll(quotedprintable.NewReader(bytes.NewReader(bstr)))
	case "B":
		bstr, err = base64.StdEncoding.DecodeString(string(bstr))
	default:
		//not set encoding type

	}
	return bstr, err
}

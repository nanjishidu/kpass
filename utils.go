package main

import (
	"crypto/rand"
	r "math/rand"
	"os"
	"regexp"
	"time"
)

// https://github.com/astaxie/beego/blob/master/utils/rand.go
var alphaNum = []byte(`0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz`)

// RandomCreateBytes generate random []byte by specify chars.
func RandomCreateBytes(n int, alphabets ...byte) []byte {
	if len(alphabets) == 0 {
		alphabets = alphaNum
	}
	var bytes = make([]byte, n)
	var randBy bool
	if num, err := rand.Read(bytes); num != n || err != nil {
		r.Seed(time.Now().UnixNano())
		randBy = true
	}
	for i, b := range bytes {
		if randBy {
			bytes[i] = alphabets[r.Intn(len(alphabets))]
		} else {
			bytes[i] = alphabets[b%byte(len(alphabets))]
		}
	}
	return bytes
}

// scan dir
func ScanDumpFile(dir string, fns ...string) (fnss []string, err error) {
	f, err := os.Open(dir)
	if err != nil {
		return
	}
	defer f.Close()
	list, err := f.Readdir(-1)
	if err != nil {
		return
	}
	var fn string
	if len(fns) > 0 && fns[0] != "" {
		fn = fns[0]
	}
	for _, ff := range list {
		if ff.IsDir() || (ff.Mode()&os.ModeSymlink) > 0 {
			continue
		}
		fileName := ff.Name()
		if fn != "" {
			matched, err := regexp.MatchString(".*"+fn+".*", fileName)
			if err != nil {
				break
			}
			if !matched {
				continue
			}
		}
		matched, err := regexp.MatchString("^\\.", fileName)
		if err != nil {
			break
		}
		if matched {
			continue
		}
		fnss = append(fnss, fileName)
	}
	return

}

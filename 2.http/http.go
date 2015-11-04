package http

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
)

// GetGolang gets the front page of the golang homepage
// and to compute its sum MD5, then return a human readable version
// of the MD5
func GetGolang() (string, error) {
	res, err := http.Get("https://golang.org")
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	hash := md5.New()
	_, err = io.Copy(hash, res.Body)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

package api

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"

	// "log"
	"net/http"
	"net/http/cookiejar"
)

type Ikuai struct {
	client   *http.Client
	url      string
	username string
	passwd   string
	pass     string
}

const salt = "salt_11"

func encode(password string) string {
	hash := md5.New()
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil))
}

func NewIkuai(url, username, password string) Ikuai {
	ikuai := Ikuai{
		username: username,
	}

	ikuai.passwd = encode(password)
	ikuai.pass = base64.StdEncoding.EncodeToString([]byte(salt + password))

	if url[len(url)-1] == '/' {
		url = url[:len(url)-1]
	}

	ikuai.url = url

	cookieJar, _ := cookiejar.New(nil)
	ikuai.client = &http.Client{Jar: cookieJar}

	// log.Println(ikuai.pass)
	// log.Println(ikuai.passwd)
	// log.Println(ikuai.username)

	return ikuai
}

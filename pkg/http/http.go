package http

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"github.com/ennetech/consul-dns/pkg/config"
	"math/rand"
	"net/http"
	"time"
)

func Init(config config.ConsulDnsConfig) {
	http.HandleFunc("/", handler)
	go http.ListenAndServe(":"+config.SystemConfig.HttpPort, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// GENERATE TSIG KEYS
	generatingSecret := randStringWithCharset(32, "abcdefghijklmnopqrstuvwxyz")
	res := "Hello from consul-dns, here's some random TSIG encryption keys: (SECRET: " + generatingSecret + ")<br>"

	// hmac-md5
	hmacmd5 := hmac.New(md5.New, []byte(generatingSecret))
	res += "hmac-md5: " + base64.StdEncoding.EncodeToString(hmacmd5.Sum(nil)) + "<br>"
	// hmac-sha1
	hmacsha1 := hmac.New(sha1.New, []byte(generatingSecret))
	res += "hmac-sha1: " + base64.StdEncoding.EncodeToString(hmacsha1.Sum(nil)) + "<br>"
	// hmac-sha256
	hmacsha256 := hmac.New(sha256.New, []byte(generatingSecret))
	res += "hmac-sha256: " + base64.StdEncoding.EncodeToString(hmacsha256.Sum(nil)) + "<br>"
	// hmac-sha512
	hmacsha512 := hmac.New(sha512.New, []byte(generatingSecret))
	res += "hmac-sha512: " + base64.StdEncoding.EncodeToString(hmacsha512.Sum(nil)) + "<br>"

	fmt.Fprintf(w, "<html><body><pre>"+res+"</pre></body></html>")

}

func randStringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

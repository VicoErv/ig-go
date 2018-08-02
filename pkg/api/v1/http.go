package v1

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"crypto/tls"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"

	"github.com/go-resty/resty"
)

// ComputeHash compute hash
func ComputeHash(data string) string {
	bSigKey, _ := hex.DecodeString(SigKey)
	bData, _ := hex.DecodeString(data)

	mac := hmac.New(sha256.New, bSigKey)
	mac.Write(bData)

	exportedMAC := mac.Sum(nil)

	return hex.EncodeToString(exportedMAC)
}

// GenerateUUID generate uuid
func GenerateUUID() string {

	var format = "xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx"

	rand.Seed(time.Now().UTC().UnixNano())

	var result string

	for _, c := range format {
		if c == '-' || c == '4' {
			result += string(c)
			continue
		}

		r := int(math.Round(rand.Float64()*16.0)) | 0

		v := (r&0x3 | 0x8)

		if c == 'x' {
			v = r
		}

		result += strconv.FormatInt(int64(v), 16)
	}

	return result
}

// CreateMD5 create md5
func CreateMD5(data string) string {
	algo := md5.New()
	algo.Write([]byte("password"))

	md5 := hex.EncodeToString(algo.Sum(nil))

	return md5
}

// GenerateDeviceID Generate device id
func GenerateDeviceID() string {
	// timestamp = Time.now.to_i.to_s
	// 'android-' + create_md5(timestamp)[0..16]

	var timestamp = time.Now().UTC().UnixNano()

	return "android-" + CreateMD5(string(timestamp))[0:16]
}

// GenerateSignature generate signature
func GenerateSignature(data interface{}) string {
	// data = data.to_json
	// compute_hash(data) + '.' + data
	b, err := json.Marshal(data)

	if err != nil {
		fmt.Println(err.Error())
	}

	sB := string(b)

	return ComputeHash(string(sB)) + "." + string(sB)
}

// Http is builder
type Http struct {
	url     string
	method  string
	body    string
	headers map[string]string
}

// Get get http
func (b *Http) Get(url string) *Http {
	b.url = url
	b.method = "GET"

	return b
}

// Post is post
func (b *Http) Post(url string, body string) *Http {
	b.url = url
	b.body = body
	b.method = "POST"

	return b
}

// With is with
func (b *Http) With(headers map[string]string) *Http {
	b.headers = headers

	return b
}

// Exec execute http
func (b *Http) Exec() string {
	var result string

	if b.method == "GET" {
		var response = b.execGet()
		result = string(response.Body())
	} else if b.method == "POST" {
		var response = b.execPost()
		result = string(response.Body())
	}

	return result
}

func (b *Http) execGet() *resty.Response {
	resp, _ := resty.R().
		SetHeaders(map[string]string{
			"Accept":               "*/*",
			"Accept-Encoding":      "gzip, deflate, sdch",
			"Accept-Language":      "en-US",
			"X-IG-Capabilities":    Capabilities,
			"X-IG-Connection-Type": Type,
		}).
		SetHeaders(b.headers).
		Get(Url + b.url)

	return resp
}

func (b *Http) execPost() *resty.Response {
	resty.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

	resp, err := resty.R().
		SetHeaders(map[string]string{
			"Accept":               "*/*",
			"Accept-Encoding":      "gzip, deflate, sdch",
			"Accept-Language":      "en-US",
			"X-IG-Capabilities":    Capabilities,
			"X-IG-Connection-Type": Type,
			"Content-Type":         "application/x-www-form-urlencoded",
		}).
		SetHeaders(b.headers).
		SetBody(b.body).
		Post(Url + b.url)

	if err != nil {
		fmt.Println(err.Error())
	}

	return resp
}

package v1

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"math"
	"math/rand"
	"strconv"
	"time"
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
	b, _ := json.Marshal(data)
	sB := string(b)

	return ComputeHash(string(sB)) + "." + string(sB)
}

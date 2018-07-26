package v1

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
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

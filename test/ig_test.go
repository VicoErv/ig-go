package test

import (
	"os"
	"testing"

	"github.com/vicoerv/ig-go/pkg/api/v1"
)

func TestComputeHash(t *testing.T) {
	hash := v1.ComputeHash("asd")
	t.Log(hash)
}

func TestGenerateUUID(t *testing.T) {
	uuid := v1.GenerateUUID()
	t.Log(uuid)
}

func TestHttp(t *testing.T) {
	type Login struct {
		device_id          string
		login_attempt_user int
		password           string
		username           string
		_csrftoken         string
		_uuid              string
	}

	http := v1.Http{}

	body := v1.GenerateSignature(map[string]interface{}{
		"device_id":          v1.GenerateDeviceID(),
		"login_attempt_user": 0,
		"password":           os.Getenv("INSTAGRAM_PASSWORD"),
		"username":           os.Getenv("INSTAGRAM_USER"),
		"_csrftoken":         "missing",
		"_uuid":              v1.GenerateUUID(),
	})

	response := http.
		Post("accounts/login/",
			"ig_sig_key_version=4&signed_body="+body).
		With(map[string]string{
			"Content-Type": "application/x-www-form-urlencoded",
		}).
		Exec()
	t.Log(response)
}

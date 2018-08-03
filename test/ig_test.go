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

	response := v1.Login(os.Getenv("INSTAGRAM_USER"), os.Getenv("INSTAGRAM_PASSWORD"))

	t.Log(response)
}

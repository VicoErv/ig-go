package test

import (
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

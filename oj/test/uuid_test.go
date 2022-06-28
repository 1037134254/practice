package test

import (
	"fmt"
	"github.com/satori/go.uuid"
	"testing"
)

func TestUUid(t *testing.T) {
	uuid := uuid.NewV4().String()
	fmt.Printf("UUIDv4: %s\n", uuid)
}

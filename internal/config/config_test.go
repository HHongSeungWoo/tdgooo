package config

import (
	"testing"
)

func TestInit(t *testing.T) {
	if err := Init("../../.env"); err != nil {
		t.Error(err)
	}
}

func TestMustInit(t *testing.T) {
	MustInit("../../.env")
}

package config

import (
	"os"
	"testing"
)

func TestInit(t *testing.T) {
	if err := Init(".valid.env"); err != nil {
		t.Error(err)
	}
}

func TestInitWithInvalidPath(t *testing.T) {
	if err := Init(".invalid.path.env"); err == nil {
		t.Error("올바르지 않은 경로 처리 안됨")
	}
}

func TestInitWithInvalidData(t *testing.T) {
	origin := os.Getenv("DB_PORT")
	_ = os.Setenv("DB_PORT", "99999")
	if err := Init(".valid.env"); err == nil {
		t.Error("올바르지 않은 데이터 처리가 되지 않음")
	}
	_ = os.Setenv("DB_PORT", origin)
}

func TestInitWithInvalidType(t *testing.T) {
	origin := os.Getenv("DB_PORT")
	_ = os.Setenv("DB_PORT", "-")
	if err := Init(".valid.env"); err == nil {
		t.Error("err")
	}
	_ = os.Setenv("DB_PORT", origin)
}

func TestMustInit(t *testing.T) {
	MustInit(".valid.env")
}

func TestMustInit2(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("에러 처리 안됨")
		}
	}()
	origin := os.Getenv("DB_PORT")
	_ = os.Setenv("DB_PORT", "12123123123123")
	MustInit(".valid.env")
	_ = os.Setenv("DB_PORT", origin)
}

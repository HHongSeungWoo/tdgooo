package config

import (
	"reflect"
	"testing"
)

func TestInit(t *testing.T) {
	Init("../../.env")
	if reflect.DeepEqual(DB, &database{}) {
		t.Error("설정 셋업 실패")
	}
}

package config

import (
	"reflect"
	"testing"
)

func TestSetup(t *testing.T) {
	if reflect.DeepEqual(DBConfig, &DatabaseConfig{}) {
		t.Error("설정 셋업 실패")
	}
}

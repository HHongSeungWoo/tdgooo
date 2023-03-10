package validation

import (
	"testing"
)

type TestStruct struct {
	Test string `validate:"required,gt=3"`
}

func TestValidateStruct(t *testing.T) {
	if err := Struct(TestStruct{Test: "Test"}); err != nil {
		t.Error(err)
	}
	if err := Struct(TestStruct{}); err == nil {
		t.Error("wrong validation")
	}
}

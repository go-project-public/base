package validator_test

import (
	"testing"

	"github.com/go-project-public/base/validator"
)

type TestValidate struct {
	Name string `validate:"required"`
	Age  int64  `validate:"required"`
}

func NewTestValidate() TestValidate {
	return TestValidate{Name: "test", Age: 10}
}

func TestTestValidate(t *testing.T) {
	err := validator.Validate(NewTestValidate())
	if err != nil {
		t.Fatal(err)
	}
}

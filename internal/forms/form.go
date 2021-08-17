package forms

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap/zapcore"
	"strings"
)

type BaseForm struct{}

func (f *BaseForm) String() string {
	b, _ := json.MarshalIndent(f, "", strings.Repeat(" ", 4))
	return string(b)
}

func (f *BaseForm) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	return fmt.Errorf("NotImplementsError. ")
}

func (f *BaseForm) Unpack() error {
	return fmt.Errorf("NotImplementsError. ")
}

func (f *BaseForm) Validate() error {
	return fmt.Errorf("NotImplementsError. ")
}

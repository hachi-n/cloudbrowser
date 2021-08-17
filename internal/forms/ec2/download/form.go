package download

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/hachi-n/cloudbrowser/internal/forms"
	"go.uber.org/zap/zapcore"
)

type DownloadForm struct {
	forms.BaseForm
	Region string
}

func New(region string) *DownloadForm {
	return &DownloadForm{
		Region: region,
	}
}

func (f *DownloadForm) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("region", f.Region)
	return nil
}

func (f *DownloadForm) Unpack() string {
	return f.Region
}

func (f *DownloadForm) Validate() error {
	return validation.ValidateStruct(f,
		// Region
		validation.Field(&f.Region, validation.Length(0, 30)),
	)
}

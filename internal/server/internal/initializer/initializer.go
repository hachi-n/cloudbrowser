package initializer

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rakyll/statik/fs"
	"html/template"
	"io"
	"net/http"
	"os"
	"strings"
)

func Initialize(engine *gin.Engine) {
	loadTemplate(engine)
}

func loadTemplate(engine *gin.Engine) {
	statikFS, err := fs.New()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	walkFn := setTemplateWalkFunc(statikFS, engine)

	err = fs.Walk(statikFS, "/html", walkFn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return
}

func setTemplateWalkFunc(statikFS http.FileSystem, engine *gin.Engine) func(path string, f os.FileInfo, err error) error {
	templateFileSuffix := ".html"

	return func(path string, f os.FileInfo, err error) error {
		if f.IsDir() {
			return nil
		}

		if strings.HasSuffix(f.Name(), templateFileSuffix) {
			reader, err := statikFS.Open(path)
			if err != nil {
				return err
			}
			b, err := io.ReadAll(reader)
			if err != nil {
				return err
			}

			t, err := template.New(f.Name()).Parse(string(b))
			if err != nil {
				return err
			}
			engine.SetHTMLTemplate(t)
		}
		return nil
	}
}

package cloudbrowser_server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hachi-n/cloudbrowser/internal/handlers/ec2"
	"github.com/hachi-n/cloudbrowser/internal/server/internal/middleware"
	_ "github.com/hachi-n/cloudbrowser/pack/assets"
	"github.com/rakyll/statik/fs"
	"html/template"
	"io"
	"os"
	"strings"
)

func StartDaemon() error {
	engine := gin.Default()
	engine.Use(middleware.ServerLogFormat)
	//engine.Static("/assets", "./assets")

	initTemplate(engine)

	engine.GET("/ec2", ec2.Index)
	engine.Run(":3000")

	return nil
}

func initTemplate(engine *gin.Engine) {
	statikFS, err := fs.New()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	fn := func(path string, f os.FileInfo, err error) error {
		if f.IsDir() {
			return nil
		}

		if strings.HasSuffix(f.Name(), ".html") {
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

	err = fs.Walk(statikFS, "/html", fn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return
}

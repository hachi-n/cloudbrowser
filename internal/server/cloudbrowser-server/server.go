package cloudbrowser_server

import (
	"fmt"
	_ "github.com/hachi-n/cloudbrowser/pack"
	"github.com/markbates/pkger"
	"os"
)

func StartDaemon() error {
	initTemplate()
	//engine := gin.Default()
	//engine.Use(middleware.ServerLogFormat)
	////engine.Static("/assets", "./assets")
	//engine.LoadHTMLGlob("assets/html/ec2/*.html")
	//engine.LoadHTMLFiles()
	//
	//engine.SetHTMLTemplate()
	//
	//engine.GET("/ec2", ec2.Index)
	//engine.Run(":3000")
	//
	return nil
}

func initTemplate() {

	//	statikFS, err := fs.New()
	//	if err != nil {
	//		panic(err)
	//	}
	//
	fn := func(path string, f os.FileInfo, err error) error {
		fmt.Println(path)

		//if f.IsDir() != true && strings.HasSuffix(f.Name(), ".html") {
		//	var err error
		//	tmpl, err = tmpl.Parse()
		//	if err != nil {
		//		return err
		//	}
		//}
		//return nil
		return nil
	}
	//	_ = fn
	//
	//	f, err := statikFS.Open("/html/ec2/index.html")
	//	if err != nil {
	//		fmt.Println("assets", err)
	//	}

	err := pkger.Walk("/assets", fn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//f, err := pkger.Open("/assets/html/ec2/index.html")
	//f, err := statikFS.Open("/aws.yaml")
	//if err != nil {
	//	fmt.Println("configs", err)
	//}

	////b, err := io.ReadAll(f)
	////if err != nil {
	////	fmt.Println(err)
	////	os.Exit(1)
	////}
	//
	//fmt.Print(string(b))

	//err = fs.Walk(statikFS, "/assets", fn)
	//if err != nil {
	//	panic(err)
	//}
	return
}

//func setTemplate(router *gin.Engine) {
//	tmpl := template.New()
//
//	err := box.Walk("", fn)
//	if err != nil {
//		panic(err)
//	}
//
//	router.SetHTMLTemplate(tmpl)
//}

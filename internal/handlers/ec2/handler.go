package ec2

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func Index(c *gin.Context) {
	f, err := os.Open("data/ec2.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	b, err := io.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	output := new(ec2.DescribeInstancesOutput)
	if err := json.Unmarshal(b, output); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//for _, v := range output.Reservations {
	//	for _, i := range v.Instances {
	//		for _ , t := range i.Tags {
	//			if t.Key == "Name" {
	//
	//			}
	//		}
	//	}
	//}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"data": output,
	})
}

package main

import (
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/golang/go/src/fmt"
	"golang.org/x/net/context"

	"github.com/aws/aws-sdk-go-v2/service/athena"
)

//var client *aws.Client

func init() {
}

func main() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic(err.Error())
	}

	client := athena.New(cfg)
	req := client.StartQueryExecutionRequest()
	resp, err := req.Send(context.TODO())
	if err == nil {
		fmt.Println(resp)
	}

}

//func excecuteAthenaQuery(query *string) (id *string, err error) {
//	resultconf := &athena.ResultConfiguration{}
//
//	resultconf.
//
//}

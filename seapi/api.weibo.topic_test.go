package seapi

import (
	"context"
	"fmt"
	"github.com/wuxiaoxiaoshen/search/setransport"
	"io/ioutil"
	"testing"
)

func TestWeiBoTopic(t *testing.T) {
	req := WeiBoTopicRequest{
		Query: "杨幂",
		Host:  defaultWBHost,
	}
	client := setransport.NewClient(req.Query)
	response, _ := req.Do(context.TODO(), client)
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
}

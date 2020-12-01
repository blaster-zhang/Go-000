package main

import (
	stdErr "errors"
	"fmt"
	"github.com/pkg/errors"
)

var ErrNoRows = stdErr.New("sql.ErrNoRows")

func main() {
	xId := 123
	fmt.Printf("call api:%+v", api(xId))
}

func api(xId int) *Response {
	data, err := service(xId)
	if err != nil {
		fmt.Printf("write err log:%+v \r\n", err)
		//忽略其他错误判断
		if stdErr.Is(err, ErrNoRows) {
			return &Response{
				Code:    404,
				Message: "记录不存在",
			}
		}
	}

	//忽略vo封装
	return &Response{
		Code:    1000,
		Message: "",
		Data:    data,
	}
}

func service(xId int) (*DataEntity, error) {
	//忽略bo封装
	return dao(xId)
}

func dao(xId int) (*DataEntity, error) {
	return nil, errors.Wrap(ErrNoRows, "query empty")
}

type DataEntity struct {
	Id   int
	Name string
}

type Response struct {
	Code    int
	Data    interface{}
	Message string
}

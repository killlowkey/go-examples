package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

type Server struct{}

type ServerContext struct {
	data []byte
	code int
}

func (c *ServerContext) Write(code int, data []byte) {
	c.code = code
	c.data = data
}

func (c *ServerContext) Debug() {
	fmt.Printf("code=%d msg=%s\n", c.code, string(c.data))
}

type CreateMeetingRequest struct {
	Name string `json:"name"`
}

func (s *Server) Hello(ctx *ServerContext, req *CreateMeetingRequest) {
	fmt.Printf("I am Hello, %s\n", req.Name)
	ctx.Write(200, []byte("success"))
}

func (s *Server) World(ctx *ServerContext, req *CreateMeetingRequest) {
	fmt.Printf("I am World, %s\n", req.Name)
	ctx.Write(200, []byte("success"))
}

type MethodType struct {
	method     reflect.Value
	methodName string
	reqType    reflect.Type
}

func NewMethodType(method reflect.Value, methodName string, reqType reflect.Type) *MethodType {
	return &MethodType{
		method:     method,
		methodName: methodName,
		reqType:    reqType,
	}
}

func (m *MethodType) Call(data []byte) ([]reflect.Value, error) {
	req := reflect.New(m.reqType).Interface()
	if err := json.Unmarshal(data, &req); err != nil {
		return nil, err
	}

	ctx := &ServerContext{}
	defer ctx.Debug()

	return m.method.Call([]reflect.Value{
		reflect.ValueOf(ctx),
		reflect.ValueOf(req),
	}), nil
}

func isServerContextType(source reflect.Type) bool {
	return source == reflect.TypeOf((*ServerContext)(nil)).Elem()
}

func ScanMethod(obj any) ([]*MethodType, error) {
	objType := reflect.TypeOf(obj)
	objValue := reflect.ValueOf(obj)

	// 判断对象是否为指针
	if objType.Kind() != reflect.Pointer {
		return nil, errors.New("any is not pointer")
	}

	var res []*MethodType
	for i := 0; i < objValue.NumMethod(); i++ {
		method := objValue.Method(i)
		if !method.IsValid() {
			continue
		}

		// 校验方法参数个数是否满足2，第一个参数其实是对象指针，所以需要减去1
		if objType.Method(i).Type.NumIn()-1 != 2 {
			continue
		}

		// 校验第一个方法参数是否为 ServerContext
		if !isServerContextType(method.Type().In(0).Elem()) {
			continue
		}

		res = append(res, NewMethodType(
			method,
			objType.Method(i).Name,
			method.Type().In(1).Elem(),
		))
	}

	return res, nil
}

func main() {
	methods, err := ScanMethod(&Server{})
	if err != nil {
		panic(err)
	}

	for _, method := range methods {
		_, _ = method.Call([]byte(`{"name":"ray"}`))
	}
}

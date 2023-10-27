package main

import (
	"errors"
	"fmt"
	"log"
	"reflect"
)

type Commands struct {
	Name string
}

func (c *Commands) Cmd903(name string) (string, error) {
	log.Printf("My name is %s\n", name)
	return "My name is cmd903", nil
}

type Reflector struct {
	cache map[string]*MethodInfo
}

type MethodInfo struct {
	structType  reflect.Type
	structValue reflect.Value
	structName  string
	methodValue reflect.Value // 用于方法调用
	methodName  string        // 方法名
}

func NewReflector() *Reflector {
	return &Reflector{cache: map[string]*MethodInfo{}}
}

func (r *Reflector) Register(source any) error {
	if reflect.ValueOf(source).Kind() != reflect.Pointer {
		return errors.New("source is not a pointer")
	}

	// 通过反射获取结构体的类型，获取字段和方法信息
	structType := reflect.TypeOf(source)
	// 通过反射获取结构体的值,值用于获取值，调用方法
	structValue := reflect.ValueOf(source)

	// type 和 value 不是同一种东西
	for i := 0; i < structType.NumMethod(); i++ {
		methodType := structType.Method(i)
		methodValue := structValue.Method(i)

		// 获取 struct 名字
		structName := structType.Name()
		if len(structName) == 0 {
			structName = structType.String()
		}

		r.cache[methodType.Name] = &MethodInfo{
			structType:  structType,
			structValue: structValue,
			structName:  structName,
			methodValue: methodValue,
			methodName:  methodType.Name,
		}
		log.Printf("注册 %s struct 的 %s 方法\n", structName, methodType.Name)
	}
	return nil
}

func (r *Reflector) Call(methodName string, args []reflect.Value) ([]reflect.Value, error) {
	// 查询方法是否存在
	methodInfo, ok := r.cache[methodName]
	if !ok {
		return nil, fmt.Errorf("method %s not found", methodName)
	}

	// 调用方法
	log.Printf("调用 %s struct 的 %s 方法, 参数 %v\n", methodInfo.structName, methodInfo.methodName, args)
	return methodInfo.methodValue.Call(args), nil
}

func main() {
	commands := &Commands{Name: "Commands"}

	reflector := NewReflector()
	err := reflector.Register(commands)
	if err != nil {
		panic(err)
	}

	res, err := reflector.Call("Cmd903", []reflect.Value{reflect.ValueOf("hello")})
	if err != nil {
		panic(err)
	}

	res1 := res[0].String()
	res2 := res[1].Interface()
	log.Printf("call result1: %v\n", res1)
	log.Printf("call result2: %v\n", res2)
}

package main

import (
	"encoding/binary"
	"fmt"
	"math"
)

func encodeDecodeData() {
	// 字符串
	str := "Hello, World!"
	fmt.Println("Original string:", str)
	encodedStr := encodeString(str)
	decodedStr, length := decodeString(encodedStr)
	fmt.Println("Decoded string:", decodedStr)
	fmt.Println("Length of decoded string:", length)
	fmt.Println()

	// 整数
	x := int64(12345678)
	fmt.Println("Original integer:", x)
	encodedInt := encodeInteger(x)
	decodedInt := decodeInteger(encodedInt)
	fmt.Println("Decoded integer:", decodedInt)
	fmt.Println()

	// 浮点数
	pi := math.Pi
	fmt.Println("Original float:", pi)
	encodedFloat := encodeFloat(pi)
	decodedFloat := decodeFloat(encodedFloat)
	fmt.Println("Decoded float:", decodedFloat)
	fmt.Println()

	// 字符
	char := 'A'
	fmt.Println("Original character:", string(char))
	encodedChar := encodeChar(char)
	decodedChar := decodeChar(encodedChar)
	fmt.Println("Decoded character:", string(decodedChar))
	fmt.Println()
}

// encodeString 编码字符串
func encodeString(str string) []byte {
	buf := make([]byte, binary.MaxVarintLen64)
	n := binary.PutVarint(buf, int64(len(str)))
	buf = buf[:n]
	buf = append(buf, []byte(str)...)
	return buf
}

// decodeString 解码字符串, 返回字符串和长度
func decodeString(buf []byte) (string, int) {
	length, n := binary.Varint(buf)
	str := string(buf[n:])
	return str, int(length)
}

// encodeInteger 编码整数
func encodeInteger(x int64) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(x))
	return buf
}

// decodeInteger 解码整数
func decodeInteger(buf []byte) int64 {
	y := binary.BigEndian.Uint64(buf)
	return int64(y)
}

// encodeFloat 编码浮点数
func encodeFloat(x float64) []byte {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, math.Float64bits(x))
	return buf
}

// decodeFloat 解码浮点数
func decodeFloat(buf []byte) float64 {
	y := math.Float64frombits(binary.LittleEndian.Uint64(buf))
	return y
}

// encodeChar 编码字符
func encodeChar(char rune) []byte {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(char))
	return buf
}

// decodeChar 解码字符
func decodeChar(buf []byte) rune {
	y := binary.BigEndian.Uint32(buf)
	return rune(y)
}

// Person 结构体
type Person struct {
	Age     int32
	Height  float64
	Name    string
	Address string
}

// 编码 Person 结构体
func encodePersonBinary(p Person) []byte {
	buf := make([]byte, 0, binary.MaxVarintLen64+4+8)

	// 编码 Age
	ageBuf := make([]byte, 4)                         // 4 字节
	binary.BigEndian.PutUint32(ageBuf, uint32(p.Age)) // 大端序
	buf = append(buf, ageBuf...)

	// 编码 Height
	heightBuf := make([]byte, 8)                                         // 8 字节
	binary.LittleEndian.PutUint64(heightBuf, math.Float64bits(p.Height)) // 小端序
	buf = append(buf, heightBuf...)

	// 编码 Name
	nameBuf := encodeString(p.Name)
	buf = append(buf, nameBuf...)

	// 编码 Address
	addressBuf := encodeString(p.Address)
	buf = append(buf, addressBuf...)

	return buf
}

// 解码 Person 结构体
func decodePersonBinary(buf []byte) Person {
	var p Person

	// 解码 Age
	p.Age = int32(binary.BigEndian.Uint32(buf[:4]))
	buf = buf[4:]

	// 解码 Height
	p.Height = math.Float64frombits(binary.LittleEndian.Uint64(buf[:8]))

	// 解码 Name
	var length int
	p.Name, length = decodeString(buf[8:])

	// 解码 Address
	p.Address, _ = decodeString(buf[8+length:])
	return p
}

func encodeAndDecodePerson() {
	// 创建 Person 实例
	p := Person{
		Name:    "Alice",
		Age:     25,
		Height:  1.68,
		Address: "Shanghai, China",
	}

	// 编码 Person
	encoded := encodePersonBinary(p)
	fmt.Printf("Encoded: % x\n", encoded)

	// 解码 Person
	decoded := decodePersonBinary(encoded)
	fmt.Printf("Decoded: %+v\n", decoded)
}

// main https://pkg.go.dev/encoding/binary
func main() {
	encodeDecodeData()
	encodeAndDecodePerson()
}

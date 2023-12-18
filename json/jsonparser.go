package jsonparser

import (
	"errors"
	"strconv"
	"strings"
)

type JSONParser struct {
	data   string
	cursor int
}

func NewJSONParser(json string) *JSONParser {
	return &JSONParser{data: json}
}

func (p *JSONParser) Parse() (interface{}, error) {
	p.skipWhitespace()
	if p.cursor >= len(p.data) {
		return nil, errors.New("empty json string")
	}

	switch p.data[p.cursor] {
	case '{':
		return p.parseObject()
	case '[':
		return p.parseArray()
	default:
		return p.parseValue()
	}
}

func (p *JSONParser) parseObject() (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	p.cursor++ // skip '{'

	for {
		p.skipWhitespace()
		if p.cursor >= len(p.data) || p.data[p.cursor] == '}' {
			p.cursor++
			break
		}

		key, err := p.parseString()
		if err != nil {
			return nil, err
		}

		p.skipWhitespace()
		if p.cursor >= len(p.data) || p.data[p.cursor] != ':' {
			return nil, errors.New("expected ':' after key in object")
		}
		p.cursor++ // skip ':'

		p.skipWhitespace()
		value, err := p.Parse()
		if err != nil {
			return nil, err
		}
		obj[key] = value

		p.skipWhitespace()
		if p.cursor < len(p.data) && p.data[p.cursor] == ',' {
			p.cursor++ // skip ','
		}
	}

	return obj, nil
}

func (p *JSONParser) parseArray() ([]interface{}, error) {
	arr := make([]interface{}, 0)
	p.cursor++ // skip '['

	for {
		p.skipWhitespace()
		if p.cursor >= len(p.data) || p.data[p.cursor] == ']' {
			p.cursor++
			break
		}

		value, err := p.Parse()
		if err != nil {
			return nil, err
		}
		arr = append(arr, value)

		p.skipWhitespace()
		if p.cursor < len(p.data) && p.data[p.cursor] == ',' {
			p.cursor++ // skip ','
		}
	}

	return arr, nil
}

func (p *JSONParser) parseValue() (interface{}, error) {
	if strings.HasPrefix(p.data[p.cursor:], "null") {
		p.cursor += 4
		return nil, nil
	} else if strings.HasPrefix(p.data[p.cursor:], "true") {
		p.cursor += 4
		return true, nil
	} else if strings.HasPrefix(p.data[p.cursor:], "false") {
		p.cursor += 5
		return false, nil
	} else if p.data[p.cursor] == '"' {
		return p.parseString()
	} else {
		return p.parseNumber()
	}
}

func (p *JSONParser) parseString() (string, error) {
	if p.data[p.cursor] != '"' {
		return "", errors.New("expected string")
	}
	p.cursor++ // skip '"'

	start := p.cursor
	for p.cursor < len(p.data) && p.data[p.cursor] != '"' {
		p.cursor++
	}

	if p.cursor >= len(p.data) {
		return "", errors.New("unexpected end of string")
	}

	value := p.data[start:p.cursor]
	p.cursor++ // skip closing '"'
	return value, nil
}

func (p *JSONParser) parseNumber() (interface{}, error) {
	start := p.cursor
	for p.cursor < len(p.data) && (p.data[p.cursor] >= '0' && p.data[p.cursor] <= '9' || p.data[p.cursor] == '.' || p.data[p.cursor] == '-') {
		p.cursor++
	}

	if start == p.cursor {
		return nil, errors.New("expected number")
	}

	numStr := p.data[start:p.cursor]
	if strings.Contains(numStr, ".") {
		return strconv.ParseFloat(numStr, 64)
	} else {
		return strconv.ParseInt(numStr, 10, 64)
	}
}

func (p *JSONParser) skipWhitespace() {
	for p.cursor < len(p.data) && (p.data[p.cursor] == ' ' || p.data[p.cursor] == '\n' || p.data[p.cursor] == '\t' || p.data[p.cursor] == '\r') {
		p.cursor++
	}
}

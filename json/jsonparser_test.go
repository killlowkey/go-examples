package jsonparser

import (
	"reflect"
	"testing"
)

func TestParseEmptyObject(t *testing.T) {
	parser := NewJSONParser(`{}`)
	result, err := parser.Parse()
	if err != nil {
		t.Errorf("Parse() returned an error: %v", err)
	}
	expected := make(map[string]interface{})
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestParseSimpleObject(t *testing.T) {
	parser := NewJSONParser(`{"key": "value"}`)
	result, err := parser.Parse()
	if err != nil {
		t.Errorf("Parse() returned an error: %v", err)
	}
	expected := map[string]interface{}{"key": "value"}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestParseArray(t *testing.T) {
	parser := NewJSONParser(`[1, 2, 3]`)
	result, err := parser.Parse()
	if err != nil {
		t.Errorf("Parse() returned an error: %v", err)
	}
	expected := []interface{}{int64(1), int64(2), int64(3)}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestParseNestedStructures(t *testing.T) {
	parser := NewJSONParser(`{"nested": {"array": [1, 2, 3]}}`)
	result, err := parser.Parse()
	if err != nil {
		t.Errorf("Parse() returned an error: %v", err)
	}
	expected := map[string]interface{}{
		"nested": map[string]interface{}{
			"array": []interface{}{int64(1), int64(2), int64(3)},
		},
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestParseBooleanAndNull(t *testing.T) {
	parser := NewJSONParser(`{"bool": true, "nullValue": null}`)
	result, err := parser.Parse()
	if err != nil {
		t.Errorf("Parse() returned an error: %v", err)
	}
	expected := map[string]interface{}{
		"bool": true, "nullValue": nil,
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

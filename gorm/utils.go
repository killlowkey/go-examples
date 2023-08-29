package main

import (
	"encoding/json"
	"github.com/google/uuid"
)

func ToJsonWithIndent(v any) string {
	marshal, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return ""
	}

	return string(marshal)
}

func ToJson(v any) string {
	marshal, err := json.Marshal(v)
	if err != nil {
		return ""
	}

	return string(marshal)
}

func UUID() string {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return ""
	}
	return newUUID.String()
}

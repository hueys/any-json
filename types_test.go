package main

import (
	"encoding/json"
	"testing"
)

func TestBasicDecode(t *testing.T) {
	example := `{
		"name": "Apple",
		"quantity": 12,
		"weight": 2.7,
		"inStock": true
		}`

	var item Item
	json.Unmarshal([]byte(example), &item)

	if item.Name != "Apple" {
		t.Errorf("expected: %v, got: %v", "Apple", item.Name)
	}

	if item.Quantity != 12 {
		t.Errorf("expected: %v, got: %v", 12, item.Quantity)
	}

	if item.Weight != 2.7 {
		t.Errorf("expected: %v, got: %v", 2.7, item.Weight)
	}

	if item.InStock != true {
		t.Errorf("expected: %v, got: %v", true, item.InStock)
	}

	if item.Metadata != nil {
		t.Errorf("expected: %v, got: %v", nil, item.Metadata)
	}

	if len(item.Reviews) != 0 {
		t.Errorf("expected: %v reviews, got: %v", 0, len(item.Reviews))
	}
}

func TestDecodeMetadata(t *testing.T) {
	example := `{
		"name": "Orange",
		"metadata": {
			"countryOfOrigin": "USA",
			"type": "Navel",
			"size": "Large"
		}
		}`

	var item Item
	json.Unmarshal([]byte(example), &item)

	if m, ok := item.Metadata.(map[string]interface{}); ok {
		if m["countryOfOrigin"] != "USA" {
			t.Errorf("expected: %v, got: %v", "USA", m["countryOfOrigin"])
		}
	} else {
		t.Errorf("unexpected error, not a map[string]interface{}: %v", item.Metadata)
	}
}

func TestDecodeReviews(t *testing.T) {
	example := `{
		"name": "Orange",
		"reviews": [{"stars": 3}, {"stars": 4}, {"stars": 2, "comment": "too sour", "user": {"id": 3277465, "email": "joe@coder.com"}}]
		}`

	var item Item
	json.Unmarshal([]byte(example), &item)

	if len(item.Reviews) != 3 {
		t.Errorf("expected: %v reviews, got: %v", 3, len(item.Reviews))
	}
}

package main

type Item struct {
	Name     string        `json:"name"`
	Quantity int           `json:"quantity"`
	Weight   float32       `json:"weight"`
	InStock  bool          `json:"inStock"`
	Metadata interface{}   `json:"metadata,omitempty"`
	Reviews  []interface{} `json:"reviews,omitempty"`
}

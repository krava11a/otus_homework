package models

import "encoding/json"

type Dialog struct {
	From string `json:"from"`
	To   string `json:"to"`
	Text string `json:"text"`
}

func (d Dialog) MarshalBinary() ([]byte, error) {
	return json.Marshal(d)
}

func (d *Dialog) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &d)
}

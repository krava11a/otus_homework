package models

import "encoding/json"

type Posts struct {
	Posts []Post `json:"posts"`
}

func (p Posts) MarshalBinary() ([]byte, error) {
	return json.Marshal(p)
}

func (p *Posts) UnmarshalBinary(data []byte) error {

	return json.Unmarshal(data, &p)
}

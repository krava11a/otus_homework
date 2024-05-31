package models

import "encoding/json"

type Post struct {
	Id_post string `json:"id_post"`
	Id_user string `json:"id_user"`
	Text    string `json:"text"`
}

func (p Post) MarshalBinary() ([]byte, error) {
	return json.Marshal(p)
}

func (p *Post) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &p)
}

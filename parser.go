package addressparser

import "encoding/json"

type Location struct {
	Code     string     `json:"code"`
	Name     string     `json:"name"`
	Children []Location `json:"children"`
}

type Address struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Province string `json:"province"`
	City     string `json:"city"`
	Area     string `json:"area"`
	Detail   string `json:"detail"`
}

type AddressParser struct {
	opts Options
}

func (p *AddressParser) Parse() (*Address, error) {
	location := []*Location{}
	if err := json.Unmarshal([]byte(locationJson), &location); err != nil {
		return nil, err
	}

	return nil, nil
}

package address_parser

type Address struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Province string `json:"province"`
	City     string `json:"city"`
	Area     string `json:"area"`
	Detail   string `json:"detail"`
}

func AddressAnalyse() *Address {
	return nil
}

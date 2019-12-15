package hello

import ()

// User user type
type User struct {
	ID   int64    `json:"id,omitempty"`
	Name string   `json:"name,omitempty"`
	Addr *Address `json:"addr,omitempty"`
}

// Address address type
type Address struct {
	City   string     `json:"city,omitempty"`
	ZIP    int        `json:"zip,omitempty"`
	LatLng [2]float64 `json:"lat_lng,omitempty"`
}

var alex = User{
	ID:   0,
	Name: "",
	Addr: &Address{
		City: "",
		ZIP:  0,
		LatLng: [2]float64{
			0.0,
			0.0,
		},
	},
}

// Hoge writes a welcome string
func Hoge() string {
	return "Hello, " + alex.Name
}

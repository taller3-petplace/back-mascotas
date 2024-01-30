package model

type Veterinary struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Address  string   `json:"address"`
	Phone    string   `json:"phone"`
	Email    string   `json:"email"`
	WebSite  string   `json:"web_site"`
	IMGUrl   string   `json:"img_url"`
	City     string   `json:"city_id"`
	Location Location `json:"location"`
	Doctors  []Doctor `json:"doctors"`
	DayGuard int      `json:"day_guard"`
}

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Doctor struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

func (v Veterinary) IsZeroValue() bool {

	var zeroValue Veterinary

	result := v.ID == zeroValue.ID
	result = result && (v.Name == zeroValue.Name)
	result = result && (v.Address == zeroValue.Address)
	result = result && (v.Phone == zeroValue.Phone)
	result = result && (v.Email == zeroValue.Email)
	result = result && (v.WebSite == zeroValue.WebSite)
	result = result && (v.IMGUrl == zeroValue.IMGUrl)
	result = result && (v.Location == zeroValue.Location)

	return result
}

package entity

type SubDistrictData struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type DistrictData struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Hub  string `json:"hub,omitempty"`
}

type CityData struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ProvinceData struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Region struct {
	Province    ProvinceData    `json:"province"`
	City        CityData        `json:"city"`
	District    DistrictData    `json:"district"`
	SubDistrict SubDistrictData `json:"sub_district"`
}

type City struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

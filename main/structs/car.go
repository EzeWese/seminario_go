package structs

type Car struct {
	Brand      string `json:"brand"`
	Model      string `json:"model"`
	EngineType string `json:"engine_type,omitempty"`
	Year       int    `json:"year"`
}

func NewCar(brand string, model string, engine_type string, year int) Car {
	return Car{
		Brand:      brand,
		Model:      model,
		EngineType: engine_type,
		Year:       year,
	}
}

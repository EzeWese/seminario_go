package structs

import "errors"

type CarDatabase struct {
	Cars []Car
}

func (c *CarDatabase) SearchCar(plate string) (*Car, error) {
	for _, car := range c.Cars {
		if car.Plate == plate {
			return &car, nil
		}
	}

	return nil, errors.New("car not found")
}

/*func main() {

	car, err := c.SearchCar("ABC123")

	if err != nil {
    	fmt.Println(err)
	}
}*/

package entity

import "errors"

type Order struct {
	Id         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

func NewOrder(id string, price, tax float64) (*Order, error) {
	order := &Order{
		Id:    id,
		Price: price,
		Tax:   tax,
	}

	err := order.Validate()

	if err != nil {
		return nil, err
	}

	return order, nil
}

func (o *Order) Validate() error {
	if o.Id == "" {
		return errors.New("id is required")
	}

	if o.Price <= 0 {
		return errors.New("price must be greater than zero")
	}

	if o.Tax < 0 {
		return errors.New("tax must be greater than zero")
	}
	return nil
}

func (o *Order) CalculateFinalPrice() error {
	o.FinalPrice = o.Price + o.Tax

	err := o.Validate()

	if err != nil {
		return err
	}

	return nil
}

package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_If_I_Get_An_Error_If_Id_Is_Blank(t *testing.T) {
	order := Order{}
	assert.Error(t, order.Validate(), "invalid id")
}


func Test_If_I_Get_An_Error_If_Price_Is_Blank(t *testing.T) {
	order := Order{ID: "123"}
	assert.Error(t, order.Validate(), "invalid price")
}

func Test_If_I_Get_An_Error_If_Price_Is_Negative(t *testing.T) {
	order := Order{ID: "123", Price: -10}
	assert.Error(t, order.Validate(), "invalid price")
}


func Test_Calculate_Final_Price(t *testing.T) {
	order := Order{ID: "123", Price: 10, Tax: 1}
	order.CalculateFinalPrice()
	assert.Equal(t, order.FinalPrice, 11.0, "invalid final price")
}
package models

type Celsius float32
type Gram float64
type Kilogram float64
type Liter float64

type NameVersion struct {
    Name string
    Version int
}

func (kg *Kilogram) ToGrams() {
    return kg * 1000.
}

func (g *Gram) ToKilograms() {
    return g / 1000.
}

func (c *Celsius) ToFahrenheit() {
    return c * (9./5.) + 32.
}
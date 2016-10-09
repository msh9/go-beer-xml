package models

type Celsius float32
type Gram float64
type Kilogram float64
type Liter float64

type NameVersion struct {
    Name string `xml:"Name"`
    Version int `xml:"Version"`
}

func (kg Kilogram) ToGrams() Gram {
    return kg * 1000.
}

func (kg Kilogram) ToPounds() float64 {
    return kg / 0.45359237
}

func (g Gram) ToKilograms()  Kilogram {
    return g / 1000.
}

func (c Celsius) ToFahrenheit() float64 {
    return c * (9./5.) + 32.
}
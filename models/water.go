package models

type Water struct {
    NameVersion
    Amount Liter
    Calcium float64
    Bicarbonate float64
    Sulfate float64
    Chloride float64
    Sodium float64
    Magnesium float64
    PH float64
    Notes string
}

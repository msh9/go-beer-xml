package models

type YeastType int
type YeastForm int
type YeastFlocculation int

const (
    Ale YeastType = iota
    Lager
    Wheat
    Wine
    Champagne
)

const (
    Liquid YeastForm = iota
    Dry
    Slant
    Cluture
)

const (
    Low YeastFlocculation = iota
    Medium
    High
    VeryHigh
)

type Yeast struct {
    NameVersion
    Type []YeastType
    Form []YeastForm
    Amount float32
    AmountIsWeight bool
    Laboratory string
    ProductID string
    MinTemperature Celsius
    MaxTemperature Celsius
    Flocculation []YeastFlocculation
    Attenuation float32
    Notes string
    BestFor string
    TimesCultured int
    MaxReuse int
    AddToSecondary bool
}

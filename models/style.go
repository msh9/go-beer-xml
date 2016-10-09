package models

type StyleType int

const (
    Lager StyleType = iota
    Ale
    Mead
    Wheat
    Mixed
    Cider
)

type Style struct {
    NameVersion
    Category string
    CategoryNumber string
    StyleLetter string
    StyleGuide string
    Type []StyleType
    OGMin float64
    OGMax float64
    FGMin float64
    FGMax float64
    IBUMin float64
    IBUMax float64
    ColorMin float64
    ColorMax float64
    CarbMin float64
    CarbMax float64
    ABVMin float64
    ABVMax float64
    Notes string
    Profile string
    Ingredients string
    Examples string
}

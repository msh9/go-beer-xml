package models

type FermentableType int

const (
    Grain FermentableType = iota
    Sugar
    Extract
    Adjunct
)

type Fermentable struct {
    NameVersion
    Type FermentableType
    Amount float32
    Yield float32
    Color float32
    AddAfterBoil bool
    Origin string
    Supplier string
    Notes string
    CoarseFineDiff float32
    Moisture float32
    DiastaticPower float32
    Protein float32
    MaxInBatch float32
    RecommendMash bool
    IbuGalPerLB float32
}

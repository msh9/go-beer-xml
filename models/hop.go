package models

type HopUse int
type HopType int
type HopForm int

const (
    Boil HopUse = iota
    DryUse
    Mash
    FirstWort
    Aroma
)

const (
    Bittering HopType = iota
    Aromatic
    Both
)

const (
    Pellet HopForm = iota
    Plug
    Leaf
)

// A Hop represents the definition of a single hop type
type Hop struct {
    Name string
    Version int
    Alpha float32
    Amount float32
    Use []HopUse
    Time int
    Notes string
    Type HopType
    Form HopForm
    Beta float32
    HSI float32
    Origin string
    Substitutes string
    Humulene float32
    Caryophyllene float32
    Cohumulene float32
    Myrcene float32
}

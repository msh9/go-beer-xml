package models

import (
    "time"
)

type MiscType int
type MiscUse int

const (
    Spice MiscType = iota
    Fining
    WaterAgent
    Herb
    Flavor
    Other
)

const (
    Boil MiscUse = iota
    Mash
    Primary
    Secondary
    Bottling
)

type Misc struct {
    NameVersion
    Type MiscType
    Use MiscUse
    Time time.Duration
    Amount float64
    AmountIsWeight bool
    UseFor string
    Notes string
}

package models

import (
    "time"
)

type MiscType int

const (
    Spice MiscType = iota
    Fining
    WaterAgent
    Herb
    Flavor
    Other
)

type Misc struct {
    NameVersion
    Type MiscType
    Use MiscUse
    Time time.Duration
    

}

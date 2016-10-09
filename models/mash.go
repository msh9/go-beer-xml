package models

import (
    "time"
)

type MashType int

const (
    Infusion MashType = iota
    Temperature
    Decoction
)

type MashStep struct {
    NameVersion
    Type []MashType
    InfuseAmount Liter
    StepTemp Celsius
    StepTime time.Duration
    RampTime time.Duration
    EndTemp Celsius
}


type MashProfile struct {
    NameVersion
    GrainTemp Celsius
    MashSteps []MashStep
    Notes string
    TunTemp Celsius
    SpargeTemp Celsius
    PH float64
    TunWeight Kilogram
    TunSpecificHeat float64
    EquipAdjust bool
}

package models

type Equipment struct {
    NameVersion
    BoilSize Liter
    BatchSize Liter
    TunVolume Liter
    TunWeight Kilogram
    TunSpecificHeat float64
    TopUpWater Liter
    TrubChillerLoss Liter
    EvapRate float64
    CalcBoilSize bool
    LauterDeadspace Liter
    TopUpKettle Liter
    HopUtilization float64
    Notes string
}

package models

import "time"

type RecipeType int

const (
    Extract RecipeType = iota
    PartialMash
    AllGrain
)

//A single beer recipe
type Recipe struct {
    NameVersion
    Type RecipeType
    Equipment Equipment
    Brewer string
    AsstBrewer string
    BatchSize Liter
    BoilSize Liter
    BoilTime time.Duration
    Efficiency float64
    Hops []Hop
    Fermentables []Fermentable
    Miscs []Misc
    Yeasts []Yeast
    Waters []Water
    Mash MashProfile
    Notes string
    TasteNotes string
    TasteRatings string
    OG float64
    FG float64
    FermentationStages int
    PrimaryAge time.Duration
    PrimaryTemp Celsius
    SecondaryAge time.Duration
    SecondaryTemp Celsius
    TertiaryAge time.Duration
    TertiaryTemp Celsius
    Age time.Duration
    AgeTemp Celsius
    Date string
    Carbonation float64
    ForcedCarbonation bool
    PrimingSugarName string
    CarbonationTemp Celsius
    PrimingSugarEquiv float64
    KegPrimingFactor float64
}

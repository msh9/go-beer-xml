package models

type FermentableType int

const (
    Grain FermentableType = iota
)

type Fermentable struct {
    Name string
    Version int
    Type FermentableType
}

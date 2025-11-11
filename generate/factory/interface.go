package factory

import "github.com/TerraDharitri/drt-go-chain-deploy/data"

// DataGenerator represents a structure that can generate genesis data
type DataGenerator interface {
	Generate() (*data.GeneratorOutput, error)
	IsInterfaceNil() bool
}

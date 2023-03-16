package data_generator

import (
	"mongo_data/constants"
	"time"
)

type Generator interface {
	GetInt32() int32
	GetInt64() int64
	GetBool() bool
	GetDate() time.Time
	GetString() string
	GetDouble() float64
}

type GeneratorType = int

var generatorMap map[GeneratorType]Generator

func init() {
	generatorMap = make(map[GeneratorType]Generator, 0)
	generatorMap[constants.Random] = &RandomGenerator{}
	generatorMap[constants.WordLib] = &WordLibGenerator{}
}

func GetGenerator(t GeneratorType) Generator {
	return generatorMap[t]
}

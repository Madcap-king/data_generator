package data_generator

import (
	"github.com/brianvoe/gofakeit/v6"
	"time"
)

type RandomGenerator struct {
}

func (p *RandomGenerator) GetInt32() int32 {
	return gofakeit.Int32()
}

func (p *RandomGenerator) GetInt64() int64 {
	return gofakeit.Int64()
}

func (p *RandomGenerator) GetBool() bool {
	return gofakeit.Bool()
}

func (p *RandomGenerator) GetDate() time.Time {
	return gofakeit.Date()
}

func (p *RandomGenerator) GetDouble() float64 {
	return gofakeit.Float64()
}

func (p *RandomGenerator) GetString() string {
	return gofakeit.Name()
}

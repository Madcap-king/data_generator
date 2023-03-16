package data_generator

import (
	"time"
)

type WordLibGenerator struct {
}

//todo

func (p *WordLibGenerator) GetInt32() int32 {
	return 0
}

func (p *WordLibGenerator) GetInt64() int64 {
	return 0
}

func (p *WordLibGenerator) GetBool() bool {
	return false
}

func (p *WordLibGenerator) GetDate() time.Time {
	return time.Now()
}

func (p *WordLibGenerator) GetDouble() float64 {
	return 0
}

func (p *WordLibGenerator) GetString() string {
	return ""
}

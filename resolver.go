package main

import (
	"context"
)

type Resolver struct {}

type InstrumentResolver struct {
	m Instrument
}

type InstrumentInput struct {
	Name string
	Digits int
}

type Instrument struct {
	Name string `db:"name"`
	Digits int `db:"digits"`
}

func (r *Resolver) GetAll() (*[]*InstrumentResolver) {

	s := InstrumentResolver{
		m:  Instrument{Name: "XXX", Digits: 3,},
	}

	l := make([]*InstrumentResolver, 1)
	l[0] = &s
	return &l
}

func (r *Resolver) UpdateInstrument(ctx context.Context, args struct{ Instrument InstrumentInput }) *InstrumentResolver {
	
	var ins = Instrument{Name: "XXX", Digits: 3}

	return &InstrumentResolver{ins}
}
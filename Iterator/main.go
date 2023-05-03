package main

import "fmt"

type FlatNumerable interface {
	CreateIterator() FlatIterator
	Count() int
	Flats() map[int]Flat
}

type PanelHouse struct {
	m map[int]Flat
}

func (p *PanelHouse) CreateIterator() FlatIterator {
	return &MapToListFlatIterator{
		numerable: p,
		i:         0,
	}
}

func (p *PanelHouse) Count() int {
	return len(p.m)
}

func (p *PanelHouse) Flats() map[int]Flat {
	return p.m
}

type FlatIterator interface {
	HasNext() bool
	Next() Flat
}

type MapToListFlatIterator struct {
	numerable FlatNumerable
	i         int
	list      []Flat
}

func (m *MapToListFlatIterator) HasNext() bool {
	return m.i < m.numerable.Count()
}

func (m *MapToListFlatIterator) Next() Flat {
	if len(m.list) == 0 {
		for _, flat := range m.numerable.Flats() {
			m.list = append(m.list, flat)
		}
	}

	res := m.list[m.i]
	m.i++

	return res
}

type Flat struct {
	Size       float32
	RoomsCount int
}

func main() {
	house := &PanelHouse{m: map[int]Flat{
		1: {
			Size:       22,
			RoomsCount: 1,
		},
		32: {
			Size:       56,
			RoomsCount: 2,
		},
	},
	}

	it := house.CreateIterator()

	for it.HasNext() {
		fmt.Printf("%+v \n", it.Next())
	}
}

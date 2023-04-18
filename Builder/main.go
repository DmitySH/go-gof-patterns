package main

import (
	"fmt"
	"time"
)

type BuildingCompany struct {
	Name     string
	Builder  HouseBuilder
	houseCnt int
}

func (b *BuildingCompany) CreateHouse() *House {
	b.houseCnt += 1
	fmt.Println("Build house number", b.houseCnt)

	return b.Builder.BuildHouse()
}

type House struct {
	Roof       string
	Walls      string
	Foundation string
}

type HouseBuilder interface {
	BuildFoundation()
	BuildWalls()
	BuildRoof()
	BuildHouse() *House
}

type SmallHouseBuilder struct {
	Material string
	house    House
}

func (s *SmallHouseBuilder) BuildFoundation() {
	s.house = House{Foundation: "stone foundation"}
}
func (s *SmallHouseBuilder) BuildWalls() {
	s.house.Walls = s.Material + "walls"
}
func (s *SmallHouseBuilder) BuildRoof() {
	s.house.Roof = "little" + s.Material + "roof"
}

func (s *SmallHouseBuilder) BuildHouse() *House {
	s.BuildFoundation()
	s.BuildWalls()
	s.BuildRoof()

	return &s.house
}

type SkyscraperBuilder struct {
	Floors int
	house  House
}

func (s *SkyscraperBuilder) BuildFoundation() {
	s.house = House{Foundation: "great foundation"}
}
func (s *SkyscraperBuilder) BuildWalls() {
	s.house.Walls = fmt.Sprintf("%d, %s!", s.Floors, "floors")
}
func (s *SkyscraperBuilder) BuildRoof() {
	s.house.Roof = "ugly plain roof"
}

func (s *SkyscraperBuilder) BuildHouse() *House {
	fmt.Println("a lot of people building a new masterpiece")

	time.Sleep(time.Second)
	s.BuildFoundation()
	time.Sleep(time.Second)
	s.BuildWalls()
	time.Sleep(time.Second)
	s.BuildRoof()

	return &s.house
}

func main() {
	var b HouseBuilder

	b = &SmallHouseBuilder{Material: "wood"}

	c := BuildingCompany{
		Name:     "TRUMP GUYS",
		Builder:  nil,
		houseCnt: 0,
	}

	c.Builder = b

	var house *House
	house = c.CreateHouse()

	fmt.Println(house)

	c.Builder = &SkyscraperBuilder{Floors: 78}

	house = c.CreateHouse()
	fmt.Println(house)
}

func (h *House) String() string {
	return fmt.Sprintf("House: %s, %s, %s", h.Foundation, h.Walls, h.Roof)
}

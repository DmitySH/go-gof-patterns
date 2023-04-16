package main

import "fmt"

type Enemy interface {
	Attack()
	Block()
}

type Archer struct {
}

func (s *Archer) Attack() {
	fmt.Println("Shooting from bow")
}

func (s *Archer) Block() {
	fmt.Println("Running away")
}

type Warrior struct {
}

func (s *Warrior) Attack() {
	fmt.Println("Sword hit")
}

func (s *Warrior) Block() {
	fmt.Println("Using shield")
}

type EnemySpawner interface {
	Spawn() Enemy
}

type ArcherSpawner struct {
}

func (s *ArcherSpawner) Spawn() Enemy {
	return &Archer{}
}

type WarriorSpawner struct {
}

func (s *WarriorSpawner) Spawn() Enemy {
	return &Warrior{}
}

func main() {
	var s EnemySpawner
	var en Enemy

	s = &ArcherSpawner{}
	en = s.Spawn()

	en.Attack()
	en.Block()

	s = &WarriorSpawner{}
	en = s.Spawn()

	en.Attack()
	en.Block()
}

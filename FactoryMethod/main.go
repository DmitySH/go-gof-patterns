package main

import "fmt"

type Enemy interface {
	Attack()
}

type EnemySpawner interface {
	Spawn() Enemy
}

type Archer struct {
}

type Warrior struct {
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

	s = &WarriorSpawner{}
	en = s.Spawn()

	en.Attack()
}

func (s *Warrior) Attack() {
	fmt.Println("Sword hit")
}

func (s *Archer) Attack() {
	fmt.Println("Shooting from bow")
}

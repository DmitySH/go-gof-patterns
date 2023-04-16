package main

import "fmt"

type Support interface {
	Heal()
}

type Enemy interface {
	Attack()
}

type NPCFactory interface {
	CreateSupport() Support
	CreateEnemy() Enemy
}

type Angel struct {
}

type Shooter struct {
}

type Doctor struct {
}

type Warrior struct {
}

type RangeNPCFactory struct {
}

func (r *RangeNPCFactory) CreateSupport() Support {
	return &Angel{}
}

func (r *RangeNPCFactory) CreateEnemy() Enemy {
	return &Shooter{}
}

type MeleeNPCFactory struct {
}

func (m *MeleeNPCFactory) CreateSupport() Support {
	return &Doctor{}
}

func (m *MeleeNPCFactory) CreateEnemy() Enemy {
	return &Warrior{}
}

func main() {
	var f NPCFactory
	var s Support
	var e Enemy

	f = &RangeNPCFactory{}
	s = f.CreateSupport()
	e = f.CreateEnemy()
	s.Heal()
	e.Attack()

	f = &MeleeNPCFactory{}
	s = f.CreateSupport()
	e = f.CreateEnemy()
	s.Heal()
	e.Attack()
}

func (a *Angel) Heal() {
	fmt.Println("Use healing magic")
}

func (s *Shooter) Attack() {
	fmt.Println("Shot you")
}

func (d *Doctor) Heal() {
	fmt.Println("Give you healing potion")
}

func (w *Warrior) Attack() {
	fmt.Println("Hit you by the sword")
}

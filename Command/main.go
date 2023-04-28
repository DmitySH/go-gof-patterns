package main

import (
	"fmt"
)

type AuraSpeaker struct {
}

func (a *AuraSpeaker) Music(volume int) {
	for i := 0; i < volume; i++ {
		fmt.Println("tuc tuc tuc tuc great music")
	}
}

func (a *AuraSpeaker) Stop() {
	fmt.Println("no more music :/")
}

type LedLight struct {
}

func (l *LedLight) On() {
	fmt.Println("lights turned on")
}

func (l *LedLight) Off() {
	fmt.Println("lights turned off")
}

type Command interface {
	Execute()
	Undo()
}

type PlayMusicCommand struct {
	Volume int
	aura   AuraSpeaker
}

func (p *PlayMusicCommand) Execute() {
	p.aura.Music(p.Volume)
}

func (p *PlayMusicCommand) Undo() {
	p.aura.Stop()
}

type TurnTheLightsCommand struct {
	lights LedLight
}

func (t *TurnTheLightsCommand) Execute() {
	t.lights.On()
}

func (t *TurnTheLightsCommand) Undo() {
	t.lights.Off()
}

type VoiceCommands map[string]Command

type Alice struct {
	Commands VoiceCommands
}

func (a *Alice) PerformCommand(cmd string) {
	a.Commands[cmd].Execute()
}

func (a *Alice) UndoCommand(cmd string) {
	a.Commands[cmd].Undo()
}

func main() {
	alice := &Alice{Commands: make(VoiceCommands)}
	alice.Commands["lets make disco"] = &PlayMusicCommand{aura: AuraSpeaker{}, Volume: 2}
	alice.Commands["turn on light"] = &TurnTheLightsCommand{lights: LedLight{}}

	alice.PerformCommand("turn on light")
	alice.PerformCommand("lets make disco")
	alice.UndoCommand("lets make disco")
	alice.UndoCommand("turn on light")
}

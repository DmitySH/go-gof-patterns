package main

import "fmt"

type Message struct {
	Content string
}

type Colleague interface {
	Send(msg Message)
	Notify(msg Message)
}

type Head struct {
	Lead Mediator
}

func (h *Head) Notify(_ Message) {

}

func (h *Head) Send(msg Message) {
	if h.Lead != nil {
		h.Lead.Send(msg, h)
	}
}

type Junior struct {
	Lead Mediator
}

func (j *Junior) Notify(msg Message) {
	fmt.Println("got new task:", msg.Content)
}

func (j *Junior) Send(msg Message) {
	if j.Lead != nil {
		j.Lead.Send(msg, j)
	}
}

type Middle struct {
	Lead Mediator
}

func (m *Middle) Notify(msg Message) {
	fmt.Println("middle done task:", msg.Content, "and ready to do new task")
}

func (m *Middle) Send(msg Message) {
	if m.Lead != nil {
		m.Lead.Send(msg, m)
	}
}

type Senior struct {
	Lead Mediator
}

func (s *Senior) Notify(msg Message) {
	fmt.Println("reviewing task:", msg.Content)
}

func (s *Senior) Send(msg Message) {
	if s.Lead != nil {
		s.Lead.Send(msg, s)
	}
}

type Mediator interface {
	Send(msg Message, colleague Colleague)
}

type TeamLead struct {
	Junior    Colleague
	Middle    Colleague
	Senior    Colleague
	Head      Colleague
	TasksDone int
}

func (t *TeamLead) Send(msg Message, colleague Colleague) {
	switch colleague {
	case t.Junior:
		t.Senior.Notify(msg)
	case t.Middle:
		t.Middle.Notify(msg)
	case t.Senior:
		fmt.Println("task:", msg.Content, "done review")
	case t.Head:
		t.Junior.Notify(msg)
	}
}

func main() {
	lead := &TeamLead{}

	jun := &Junior{Lead: lead}
	mid := &Middle{Lead: lead}
	sen := &Senior{Lead: lead}
	head := &Head{Lead: lead}

	lead.Junior = jun
	lead.Middle = mid
	lead.Senior = sen
	lead.Head = head

	jun.Send(Message{Content: "some task"})
	mid.Send(Message{Content: "other task"})
	sen.Send(Message{Content: "some task"})
	head.Send(Message{Content: "task for current sprint"})
}

package main

import "fmt"

type JiraTicketState interface {
	Forward(ticket *JiraTicket)
	Back(ticket *JiraTicket)
}

type Backlog struct {
}

func (b *Backlog) String() string {
	return "Backlog"
}

func (b *Backlog) Forward(ticket *JiraTicket) {
	ticket.state = &Todo{}
}

func (b *Backlog) Back(_ *JiraTicket) {
	fmt.Println("ticket is already in backlog")
}

type Todo struct {
}

func (t *Todo) String() string {
	return "TO DO"
}

func (t *Todo) Forward(ticket *JiraTicket) {
	ticket.state = &InProgress{}
}

func (t *Todo) Back(ticket *JiraTicket) {
	ticket.state = &Backlog{}
}

type InProgress struct {
}

func (i *InProgress) String() string {
	return "In progress"
}

func (i *InProgress) Forward(ticket *JiraTicket) {
	if ticket.minApprovesCount != 0 {
		ticket.state = &InReview{
			NeedApproves:  ticket.minApprovesCount,
			ApprovesGiven: 0,
		}
	} else {
		ticket.state = &Closed{}
	}
}

func (i *InProgress) Back(ticket *JiraTicket) {
	ticket.state = &Todo{}
}

type InReview struct {
	NeedApproves  int
	ApprovesGiven int
}

func (i *InReview) String() string {
	return fmt.Sprintf("In review. Approves: %d / %d", i.ApprovesGiven, i.NeedApproves)
}

func (i *InReview) Forward(ticket *JiraTicket) {
	if i.ApprovesGiven < i.NeedApproves {
		fmt.Printf("Need %d more approves\n", i.NeedApproves-i.ApprovesGiven)
	} else {
		ticket.state = &Closed{}
	}
}

func (i *InReview) Back(ticket *JiraTicket) {
	ticket.state = &InProgress{}
}

type Closed struct {
}

func (c *Closed) Forward(_ *JiraTicket) {
	fmt.Println("ticket is already closed")
}

func (c *Closed) Back(ticket *JiraTicket) {
	ticket.state = &Todo{}
}

func (c *Closed) String() string {
	return "Closed"
}

type JiraTicket struct {
	ID               int64
	state            JiraTicketState
	minApprovesCount int
}

func (j *JiraTicket) String() string {
	return fmt.Sprintf("Ticket #%d: %s", j.ID, j.state)
}

func (j *JiraTicket) Forward() {
	j.state.Forward(j)
}

func (j *JiraTicket) Back() {
	j.state.Back(j)
}

func (j *JiraTicket) Approve() {
	inReview, ok := j.state.(*InReview)
	if !ok {
		fmt.Printf("Tiket #%d is in %s state, not in Review\n", j.ID, j.state)
	} else {
		inReview.ApprovesGiven++
	}
}

func main() {
	ticket := &JiraTicket{ID: 1232, state: &Backlog{}, minApprovesCount: 2}
	fmt.Println(ticket)

	ticket.Back()
	ticket.Forward()
	fmt.Println(ticket)

	ticket.Forward()
	fmt.Println(ticket)

	ticket.Forward()
	fmt.Println(ticket)
	ticket.Forward()

	ticket.Approve()
	ticket.Approve()
	ticket.Forward()
	fmt.Println(ticket)

	ticket.Back()
	fmt.Println(ticket)

	ticket = &JiraTicket{ID: 1232, state: &Backlog{}, minApprovesCount: 0}

	ticket.Forward()
	ticket.Forward()
	ticket.Approve()
	fmt.Println(ticket)

	ticket.Forward()
	ticket.Forward()
	fmt.Println(ticket)
}

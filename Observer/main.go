package main

import "fmt"

type PostObserver interface {
	Update(post Post)
}

type PersonalEmail struct {
	User string
}

func (p *PersonalEmail) Update(post Post) {
	fmt.Printf("%s received new post: %s\n", p.User, post)
}

type CorporateEmail struct {
	Users []string
}

func (c *CorporateEmail) Update(post Post) {
	for _, user := range c.Users {
		fmt.Printf("%s received new post: %s\n", user, post)
	}
}

type Observable interface {
	AddObserver(observer PostObserver)
	RemoveObserver(observer PostObserver)
	NotifyObservers(post Post)
}

type Post string

type DailyBugle struct {
	subscribers []PostObserver
}

func (d *DailyBugle) AddObserver(observer PostObserver) {
	d.subscribers = append(d.subscribers, observer)
}

func (d *DailyBugle) RemoveObserver(observer PostObserver) {
	removeIdx := -1
	for i, sub := range d.subscribers {
		if sub == observer {
			removeIdx = i
			break
		}
	}

	if removeIdx != -1 {
		d.subscribers = append(d.subscribers[:removeIdx], d.subscribers[removeIdx+1:]...)
	}
}

func (d *DailyBugle) NotifyObservers(post Post) {
	for _, sub := range d.subscribers {
		sub.Update(post)
	}
}

func main() {
	personal := &PersonalEmail{User: "Alex"}
	corporate := &CorporateEmail{Users: []string{"Dimon", "Vanyok", "Ave"}}

	bugle := DailyBugle{}
	bugle.AddObserver(personal)
	bugle.AddObserver(corporate)

	bugle.NotifyObservers("HOT DAILY BUGLE! POST!")

	bugle.RemoveObserver(corporate)
	bugle.NotifyObservers("something special ^:^")
}

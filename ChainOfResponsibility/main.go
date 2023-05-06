package main

import "fmt"

type Data struct {
	DataSizeGB float64
	Permanent  bool
	Content    string
}

type DataSaver interface {
	Save(data Data)
}

type TemporarySaver struct {
	Successor DataSaver
}

func (t *TemporarySaver) Save(data Data) {
	if !data.Permanent {
		fmt.Printf("data {%s} saved to tmp memory", data.Content)
		return
	}

	if t.Successor != nil {
		fmt.Println("delegating to next successor")
		t.Successor.Save(data)
	}
}

type VirtualSaver struct {
	FreeVirtualGB float64
	Successor     DataSaver
}

func (v *VirtualSaver) Save(data Data) {
	if data.DataSizeGB <= v.FreeVirtualGB {
		fmt.Printf("data {%s} saved to virtual memory", data.Content)
		return
	}

	if v.Successor != nil {
		fmt.Println("delegating to next successor")
		v.Successor.Save(data)
	}
}

type Vault struct {
	Successor DataSaver
}

func (v *Vault) Save(data Data) {
	fmt.Printf("data {%s} saved to huge slow vault", data.Content)
}

func main() {
	var s1, s2, s3 DataSaver

	s3 = &Vault{Successor: nil}
	s2 = &VirtualSaver{Successor: s3, FreeVirtualGB: 1.93}
	s1 = &TemporarySaver{Successor: s2}

	d := Data{
		DataSizeGB: 3.412,
		Permanent:  true,
		Content:    "some cool content",
	}

	s1.Save(d)
}

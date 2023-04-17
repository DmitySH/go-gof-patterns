package singletonwarehouse

import (
	"fmt"
	"sync"
)

var warehouseInstance *Warehouse = nil
var whMu = sync.Mutex{}

type Warehouse struct {
	Name     string
	Location string
}

func (w *Warehouse) String() string {
	return fmt.Sprintf("%s: %s", w.Name, w.Location)
}

func GetWarehouse(name, location string) *Warehouse {
	whMu.Lock()
	defer whMu.Unlock()

	if warehouseInstance == nil {
		warehouseInstance = &Warehouse{
			Name:     name,
			Location: location,
		}
	}

	return warehouseInstance
}

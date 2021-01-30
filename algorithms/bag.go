package algorithms

import (
	"../structures"
)

// Bag implements a bag data type by structures
type Bag struct {
	*structures.LinkList
}

// NewBag ...
func NewBag() *Bag {
	return &Bag{structures.NewLinkList()}
}

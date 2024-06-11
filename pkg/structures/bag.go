package structures

// Bag implements a bag data type by structures
type Bag struct {
	*LinkList
}

// NewBag ...
func NewBag() *Bag {
	return &Bag{NewLinkList()}
}

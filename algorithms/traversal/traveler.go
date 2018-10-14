package traversal

// Traveler is interface of traversal structure.
type Traveler interface {
	Next() bool
	Value() interface{}
}

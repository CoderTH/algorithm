package Set

type Set interface {
	Add(e int)
	Remove(e int)
	Contains(e int) bool
	GetSize() int
	IsEmpty() bool
}

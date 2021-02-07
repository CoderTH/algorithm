package Map

type Map interface {
	Add(key interface{}, value interface{})
	Remove(key interface{}) interface{}
	Contains(key interface{}) bool
	Get(Key interface{}) interface{}
	Set(key interface{}, value interface{})
	GetSize() int
	IsEmpty() bool
}

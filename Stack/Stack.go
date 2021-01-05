package Stack

type Stack interface {
	GetSize()int
	IsEmpty()bool
	Push(e interface{})
	Pop()interface{}
	Peek()interface{}
	String()string
}

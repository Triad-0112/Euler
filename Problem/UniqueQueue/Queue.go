package queue

type Queue interface {
	Push(key interface{})
	Pop() interface{}
	Contains(key interface{}) bool
	Len() int
	Keys() []interface{}
}

// Push values
func (q *Data) Push(n interface{}) *Data {
	if q.Len() < q.size {
		q.data = append(q.data, n)
	} else {
		q.Pop()
		q.Push(n)
	}
	return q
}

//Pop
func (q *Data) Pop() interface{} {
	if len(q.data) == 0 {
		return 0
	} else {
		element := q.data[0]
		q.data = q.data[1:]
		return element
	}
}

//Check the queue logic condition if it was fullfilled or no by checking the existence of item
func (q *Data) Contains(key interface{}) bool {
	cond := false
	for i := 0; i < q.Len(); i++ {
		if q.data[i] == key {
			cond = true
		}
	}
	return cond
}

//Len
func (q *Data) Len() int {
	return len(q.data)
}

type Data struct {
	size int
	data []interface{}
}

func New(size int) *Data {
	return &Data{
		size: size,
	}
}

//Keys, Show current values with fixed size
func (q *Data) Keys() []interface{} {
	return q.data
}

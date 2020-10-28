package utils

type Counter struct {
	ID int
}

func NewCounter() *Counter {
	return &Counter{
		ID: 0,
	}
}

func (c *Counter) Increment() {
	c.ID++
}

func (c *Counter) GetID() int {
	c.Increment()
	return c.ID
}

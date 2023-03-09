package pandora

type Calc struct {
	Accumulator int
}

func (c *Calc) Result() int {
	return c.Accumulator
}

func (c *Calc) Clear() *Calc {
	c.Accumulator = 0
	return c
}

func (c *Calc) Add(operand int) *Calc {
	c.Accumulator += operand
	return c
}

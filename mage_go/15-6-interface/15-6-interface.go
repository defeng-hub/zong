package main

type Transporter interface {
	move(s string, desc string) (int, error)

	sound(int) error
}

type car struct {
}

func (c *car) move(s string, desc string) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (c *car) sound(i int) error {
	//TODO implement me
	panic("implement me")
}

func main() {

}

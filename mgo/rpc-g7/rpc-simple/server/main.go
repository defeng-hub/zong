package server

type HelloService struct {
}

func (h *HelloService) Hello(req string, reply *string) error {
	*reply = "hello:" + req
	return nil
}

func main() {

}
